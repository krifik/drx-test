package helper

import (
	"math"

	"github.com/krifik/test-drx/model"
)

type Discount struct {
	Type        string  `json:"type"`
	Value       float64 `json:"value,omitempty"`
	Condition   float64 `json:"condition,omitempty"`
	MaxDiscount float64 `json:"maxDiscount,omitempty"`
	Tiers       []Tier  `json:"tiers,omitempty"`
}

type Tier struct {
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Value float64 `json:"value"`
}

type DiscountInput struct {
	OriginalPrice float64    `json:"originalPrice"`
	Discounts     []Discount `json:"discounts"`
}

func ApplyDiscounts(input model.PriceWithDiscountsRequest) float64 {
	price := input.OriginalPrice
	totalDiscount := 0.0
	capLimit := math.MaxFloat64

	for _, d := range input.Discounts {
		switch d.Type {
		case "fixed":
			totalDiscount += d.Value

		case "percentage":
			percentDiscount := (price - totalDiscount) * (d.Value / 100)
			totalDiscount += percentDiscount

		case "conditional":
			if input.OriginalPrice > d.Condition.Amount {
				totalDiscount += d.Condition.Value
			}

		case "tiered":
			for _, tier := range d.Tiers {
				if input.OriginalPrice >= tier.Min && input.OriginalPrice <= tier.Max {
					totalDiscount += tier.Value
					break
				}
			}

		case "cap":
			capLimit = d.MaxDiscount
			totalDiscount += d.MaxDiscount
		}
	}

	if totalDiscount > capLimit {
		totalDiscount = capLimit
	}
	finalPrice := input.OriginalPrice - totalDiscount
	if finalPrice < 0 {
		finalPrice = 0
	}
	return finalPrice
}
