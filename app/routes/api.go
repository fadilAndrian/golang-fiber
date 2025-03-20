package routes

import (
	"github.com/gofiber/fiber/v2"
	"learn-fibergo/app/controllers/product_controller"
	"log"
)

func Routes() {
	app := fiber.New()

	products := app.Group("/products")
	products.Get("/", product_controller.List)
	products.Get("/:id", product_controller.Show)
	products.Post("/", product_controller.Create)
	products.Put("/:id", product_controller.Update)

	log.Fatal(app.Listen(":8080"))
}
