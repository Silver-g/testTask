package main

import (
	"fmt"
	"log"
	"net/http"
	"testTask/internal/db"
	"testTask/internal/user"
)

func main() {
	// Подключение к базе данных
	db, err := db.Connect()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}

	// Создание репозитория и хэндлера для пользователей
	userRepo := user.NewUserRepository(db)
	userHandler := &user.UserHandler{Repo: userRepo}

	// Регистрируем хэндлеры
	http.HandleFunc("/register", userHandler.RegisterUser)

	// Запуск HTTP-сервера
	fmt.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
