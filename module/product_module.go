package module

import (
	"github.com/krifik/test-drx/controller"
	"github.com/krifik/test-drx/repository"
	"github.com/krifik/test-drx/service"

	"gorm.io/gorm"
)

func NewProductModule(database *gorm.DB) controller.ProductController {
	// Setup Repository
	productRepository := repository.NewProductRepositoryImpl(database)

	// Setup Service
	productService := service.NewProductServiceImpl(productRepository)

	// Setup Controller
	productController := controller.NewProductControllerImpl(productService)

	return productController
}
