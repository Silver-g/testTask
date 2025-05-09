package post

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testTask/internal/auth"
	"testTask/internal/domain"
	"testTask/pkg/response"
)

const PostIDParam = "post_id"

func (h *PostHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, response.ErrMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	userID, err := auth.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	postID := r.URL.Query().Get(PostIDParam)
	if postID == "" {
		http.Error(w, "Не указан post_id", http.StatusBadRequest)
		return
	}
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		http.Error(w, "Invalid post_id", http.StatusBadRequest)
		return
	}

	commentsEnabled, err := h.Repo.GetCommentsEnabled(postIDInt)
	if err != nil || !commentsEnabled {
		http.Error(w, response.ErrCommentsDisabled, http.StatusForbidden)
		return
	}

	var comment domain.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, response.ErrInvalidJSON, http.StatusBadRequest)
		return
	}

	commentID, err := h.Repo.CreateComment(userID, postID, comment.Content, comment.ParentID)
	if err != nil {
		http.Error(w, response.ErrCommentCreationFailed, http.StatusInternalServerError)
		return
	}

	commentResponse := map[string]interface{}{
		"id":        commentID,
		"user_id":   userID,
		"content":   comment.Content,
		"parent_id": comment.ParentID,
	}

	response.SendJSONResponse(w, http.StatusOK, commentResponse)
}
