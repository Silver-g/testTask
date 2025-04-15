package post

import (
	"encoding/json"
	"net/http"
	"testTask/internal/auth"
	"testTask/internal/boundary"
	"testTask/internal/domain"
	"testTask/pkg/response"
)

type PostHandler struct {
	Repo *PostRepository
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, response.ErrMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	userID, err := auth.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req domain.PostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, response.ErrInvalidJSON, http.StatusBadRequest)
		return
	}

	postWithUserID := boundary.MapPostRequestToPostWithUserID(req, userID)

	postID, err := h.Repo.CreatePost(postWithUserID.UserID, postWithUserID.Title, postWithUserID.Content, postWithUserID.CommentsEnabled)
	if err != nil {
		http.Error(w, response.ErrPostCreationFailed, http.StatusInternalServerError)
		return
	}

	post := boundary.MapPostWithUserIDToPost(postWithUserID, postID)

	response.SendJSONResponse(w, http.StatusCreated, post)
}
