package service

import (
	"github.com/krifik/test-drx/model"
)

type ProductService interface {
	Add(request model.CreateProductRequest) (response model.CreateProductResponse, err error)
	FindAll() ([]model.GetProductResponse, error)
}
