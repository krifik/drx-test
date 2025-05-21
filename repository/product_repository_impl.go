package repository

import (
	"errors"

	"github.com/krifik/test-drx/config"
	"github.com/krifik/test-drx/entity"
	"github.com/krifik/test-drx/exception"
	"github.com/krifik/test-drx/model"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{DB: db}
}

func (repository *ProductRepositoryImpl) FindAll() ([]entity.Product, error) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()

	var products []entity.Product
	result := repository.DB.WithContext(ctx).Find(&products)

	if result.RowsAffected < 0 {
		return nil, errors.New("product not found")
	}

	return products, nil
}

func (repository *ProductRepositoryImpl) Add(request model.CreateProductRequest) (product entity.Product, err error) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	product = entity.Product{
		Name:        request.Name,
		Quantity:    request.Quantity,
		Price:       request.Price,
		Description: request.Description,
	}
	result := repository.DB.WithContext(ctx).Create(&product)
	exception.PanicIfNeeded(result.Error)
	return product, err
}
