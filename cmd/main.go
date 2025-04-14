package main

import (
	"fmt"
	"log"

	"testTask/internal/db"
	"testTask/internal/user"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	database, err := db.Connect()
	if err != nil {
		log.Fatal("Ошибка подключения к базе:", err)
	}
	defer database.Close()

	userRepo := user.NewRepository(database)

	// Пример создания пользователя
	newUser := &user.User{
		Username: "vasya2",
		Password: "1234",
	}
	err = userRepo.CreateUser(newUser)
	if err != nil {
		log.Fatal("Ошибка при создании пользователя:", err)
	}
	fmt.Println("Пользователь создан с ID:", newUser.ID)

	// Пример получения пользователя
	u, err := userRepo.GetUserByUsername("vasya")
	if err != nil {
		log.Fatal("Ошибка при получении пользователя:", err)
	}
	fmt.Println("Найден пользователь:", u.ID, u.Username, u.Password)
}
