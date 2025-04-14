package user

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	Repo *UserRepository // Исправили тип
}

// Хэндлер для регистрации пользователя
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	userID, err := h.Repo.CreateUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Ошибка при создании пользователя", http.StatusInternalServerError)
		return
	}

	user.ID = userID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
