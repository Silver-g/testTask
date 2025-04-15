package main

import (
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
	http.HandleFunc("/posts", postHandler.CreatePost)
	http.HandleFunc("/posts/list", postHandler.GetAllPosts)
	http.HandleFunc("/comment", postHandler.CreateComment)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
