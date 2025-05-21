package config

import (
	"github.com/krifik/test-drx/config/faker"

	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeder(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: faker.ProductFaker(db)},
	}
}
