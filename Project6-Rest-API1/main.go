package main

import (
	"Rest-API/db"
	"Rest-API/user"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal("Cannot connect to database %v", err)
	}

	repo := user.NewRepository(database)
	err = repo.Migration()
	if err != nil {
		log.Fatal(err)
	}
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	app := fiber.New()
	app.Get("/users/:id", handler.Get)
	app.Post("/users", handler.Create)

	app.Listen(":8000")
}
