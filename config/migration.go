package config

import "github.com/krifik/test-drx/entity"

type Entity struct {
	Entity interface{}
}

func RegisterEntities() []Entity {
	return []Entity{
		{Entity: entity.Product{}},
	}
}
