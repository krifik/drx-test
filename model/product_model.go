package model

import (
	"time"

	"gorm.io/gorm"
)

type CreateProductRequest struct {
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

type CreateProductResponse struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Quantity    int            `json:"quantity"`
	Price       int            `json:"price"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type GetProductResponse struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	Quantity    int            `json:"quantity"`
	Price       int            `json:"price"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type Tiers struct {
	Min   float64
	Max   float64
	Value float64
}

type Conditional struct {
	Amount float64
	Value  float64
}
type Discounts struct {
	Type        string      `json:"type"`
	Value       float64     `json:"value"`
	Tiers       []Tiers     `json:"tiers"`
	Condition   Conditional `json:"condition"`
	MaxDiscount float64     `json:"maxDiscount"`
}

type AppliedDisc struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}
type PriceWithDiscountsRequest struct {
	OriginalPrice float64     `json:"originalPrice"`
	Discounts     []Discounts `json:"discounts"`
}

type PriceWithAppliedDiscountResponse struct {
	FinalPrice  float64 `json:"finalPrice"`
	AppliedDisc []AppliedDisc
}
