package routes

import (
	"github.com/krifik/test-drx/controller"
	_ "github.com/krifik/test-drx/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Route(app *fiber.App, productController controller.ProductController) {
	app.Post("/api/products", productController.Add)
	app.Get("/api/products", productController.FindAll)
	app.Get("/api/docs/*", swagger.HandlerDefault)
}
