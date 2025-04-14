package post

import (
	"encoding/json"
	"net/http"
	"strings"

	"testTask/internal/auth"
)

type PostHandler struct {
	Repo *PostRepository
}

type PostRequest struct {
	Title           string `json:"title"`
	Content         string `json:"content"`
	CommentsEnabled bool   `json:"comments_enabled"`
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Извлечение токена из заголовка Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Требуется токен авторизации", http.StatusUnauthorized)
		return
	}

	tokenParts := strings.Split(authHeader, "Bearer ")
	if len(tokenParts) != 2 {
		http.Error(w, "Невалидный формат токена", http.StatusUnauthorized)
		return
	}

	userID, err := auth.ParseToken(tokenParts[1])
	if err != nil {
		http.Error(w, "Невалидный токен", http.StatusUnauthorized)
		return
	}

	var req PostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	postID, err := h.Repo.CreatePost(userID, req.Title, req.Content, req.CommentsEnabled)
	if err != nil {
		http.Error(w, "Ошибка при создании поста", http.StatusInternalServerError)
		return
	}

	post := Post{
		ID:              postID,
		UserID:          userID,
		Title:           req.Title,
		Content:         req.Content,
		CommentsEnabled: req.CommentsEnabled,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	posts, err := h.Repo.GetAllPosts()
	if err != nil {
		http.Error(w, "Ошибка при получении постов", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, "Ошибка при кодировании ответа", http.StatusInternalServerError)
	}
}
