package controller

import (
	"github.com/krifik/test-drx/exception"
	"github.com/krifik/test-drx/model"
	"github.com/krifik/test-drx/service"

	"github.com/gofiber/fiber/v2"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductControllerImpl(productService service.ProductService) ProductController {
	return &ProductControllerImpl{ProductService: productService}
}

func (controller *ProductControllerImpl) Add(c *fiber.Ctx) error {
	var request model.CreateProductRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	response, err := controller.ProductService.Add(request)
	exception.PanicIfNeeded(err)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
func (controller *ProductControllerImpl) FindAll(c *fiber.Ctx) error {
	responses, err := controller.ProductService.FindAll()
	exception.PanicIfNeeded(err)
	return c.JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}
