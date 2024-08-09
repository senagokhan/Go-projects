package main

import (
	"Restaurant-Management-System/database"
	"Restaurant-Management-System/middleware"
	"Restaurant-Management-System/routes"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		log.Println(c.Method(), c.Path())
		return c.Next()
	})

	app.Use(middleware.Authentication())

	routes.UserRoutes(app)
	routes.FoodRoutes(app)
	routes.MenuRoutes(app)
	routes.TableRoutes(app)
	routes.OrderRoutes(app)
	routes.OrderItemRoutes(app)
	routes.InvoiceRoutes(app)

	log.Fatal(app.Listen(":" + port))
}
