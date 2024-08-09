package routes

import (
	"Restaurant-Management-System/controllers"
	"github.com/gofiber/fiber/v2"
)

func InvoiceRoutes(app *fiber.App) {
	app.Get("/invoices", controllers.GetInvoices())
	app.Get("/invoices/:invoice_id", controllers.GetInvoince())
	app.Post("/invoices", controllers.CreateInvoince())
	app.Patch("/invoices/:invoice_id", controllers.UpdateInvoince())
}
