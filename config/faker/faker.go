package faker

import (
	"github.com/krifik/test-drx/entity"

	"gorm.io/gorm"
)

func ProductFaker(db *gorm.DB) *entity.Product {
	return &entity.Product{
		Name:        "Seblak",
		Description: "Seblack yang lebat dan berbiji",
		Quantity:    69,
		Price:       1000,
	}
}
