package repository

import (
	"github.com/krifik/test-drx/entity"
	"github.com/krifik/test-drx/model"
)

type ProductRepository interface {
	Add(request model.CreateProductRequest) (product entity.Product, err error)
	FindAll() ([]entity.Product, error)
}
