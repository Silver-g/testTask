package post

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"testTask/pkg/response"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Ошибка при загрузке .env файла")
	}
}

func getEnvInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, response.ErrMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := getEnvInt("DEFAULT_LIMIT", 100)
	offset := getEnvInt("DEFAULT_OFFSET", 0)

	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			http.Error(w, "Invalid limit", http.StatusBadRequest)
			return
		}

		if limit > 100 {
			limit = 100
		}
	}

	if offsetStr != "" {
		var err error
		offset, err = strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			http.Error(w, "Invalid offset", http.StatusBadRequest)
			return
		}
	}

	posts, err := h.Repo.GetAllPosts(limit, offset)
	if err != nil {
		http.Error(w, response.ErrPostRetrievalFailed, http.StatusInternalServerError)
		return
	}

	response.SendJSONResponse(w, http.StatusOK, posts)
}
