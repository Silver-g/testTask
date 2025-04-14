package user

import (
	"encoding/json"
	"log"
	"net/http"
	"testTask/internal/auth" // добавляем импорт пакета auth
)

type UserHandler struct {
	Repo *UserRepository
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		log.Printf("Ошибка при декодировании JSON: %v", err)
		return
	}

	userID, err := h.Repo.CreateUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Ошибка при создании пользователя", http.StatusInternalServerError)
		log.Printf("Ошибка при создании пользователя: %v", err)
		return
	}

	user.ID = userID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Хэндлер для логина и получения JWT
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var creds User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	user, err := h.Repo.GetUserByCredentials(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, "Неверное имя пользователя или пароль", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user.ID) // используем функцию из пакета auth
	if err != nil {
		http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
