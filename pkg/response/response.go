package response

import (
	"encoding/json"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Ошибка при кодировании ответа", http.StatusInternalServerError)
		return
	}
}
