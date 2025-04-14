package main

import (
	"fmt"
	"log"
	"net/http"
	"testTask/internal/db"
	"testTask/internal/post"
	"testTask/internal/user"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Не удалось загрузить .env файл, используем переменные окружения из системы")
	}
	// Подключение к базе данных
	db, err := db.Connect()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}

	userRepo := user.NewUserRepository(db)
	userHandler := &user.UserHandler{Repo: userRepo}

	http.HandleFunc("/register", userHandler.RegisterUser)
	http.HandleFunc("/login", userHandler.Login)

	postRepo := post.NewPostRepository(db)
	postHandler := &post.PostHandler{Repo: postRepo}
	http.HandleFunc("/posts", postHandler.CreatePost)       // Создание поста
	http.HandleFunc("/posts/list", postHandler.GetAllPosts) // Получение всех постов

	// Запуск HTTP-сервера
	fmt.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
