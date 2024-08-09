package routes

import (
	"Restaurant-Management-System/controllers"
	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(app *fiber.App) {
	app.Get("/orders", controllers.GetOrders())
	app.Get("/orders/:order_id", controllers.GetOrder())
	app.Post("/orders", controllers.CreateOrder())
	app.Patch("/orders/:order_id", controllers.UpdateOrder())
}
