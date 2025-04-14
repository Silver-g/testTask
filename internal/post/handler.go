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

// Добавим новый обработчик для создания комментариев
func (h *PostHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
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

	// Получение ID поста из параметров URL
	postID := r.URL.Query().Get("post_id")
	if postID == "" {
		http.Error(w, "Не указан post_id", http.StatusBadRequest)
		return
	}

	// Проверим, разрешены ли комментарии для данного поста
	var commentsEnabled bool
	err = h.Repo.DB.QueryRow("SELECT comments_enabled FROM posts WHERE id = $1", postID).Scan(&commentsEnabled)
	if err != nil || !commentsEnabled {
		http.Error(w, "Комментарии под этим постом отключены", http.StatusForbidden)
		return
	}

	// Чтение данных комментария из запроса
	var comment struct {
		Content  string `json:"content"`
		ParentID *int   `json:"parent_id,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	// Вставляем комментарий в базу данных
	commentID, err := h.Repo.CreateComment(userID, postID, comment.Content, comment.ParentID)
	if err != nil {
		http.Error(w, "Ошибка при создании комментария", http.StatusInternalServerError)
		return
	}

	// Отправляем результат
	commentResponse := map[string]interface{}{
		"id":        commentID,
		"user_id":   userID,
		"content":   comment.Content,
		"parent_id": comment.ParentID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commentResponse)
}
