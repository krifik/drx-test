package service

import (
	"github.com/krifik/test-drx/entity"
	"github.com/krifik/test-drx/model"
	"github.com/krifik/test-drx/repository"
	"github.com/krifik/test-drx/validation"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductServiceImpl(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepository}
}

func (service *ProductServiceImpl) FindAll() ([]model.GetProductResponse, error) {
	products, _ := service.ProductRepository.FindAll()
	var responses []model.GetProductResponse
	for _, product := range products {
		responses = append(responses, model.GetProductResponse{
			Id:          int(product.ID),
			Name:        product.Name,
			Quantity:    product.Quantity,
			Price:       product.Price,
			Description: product.Description,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
			DeletedAt:   product.DeletedAt,
		})
	}
	return responses, nil
}

func (service *ProductServiceImpl) Add(request model.CreateProductRequest) (response model.CreateProductResponse, err error) {
	validation.Validate(request)
	product := entity.Product{
		Name:        request.Name,
		Quantity:    request.Quantity,
		Price:       request.Price,
		Description: request.Description,
	}

	product, _ = service.ProductRepository.Add(request)
	response = model.CreateProductResponse{
		ID:          int(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Quantity:    product.Quantity,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		DeletedAt:   product.DeletedAt,
	}
	return response, err
}
