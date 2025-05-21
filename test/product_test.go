package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/krifik/test-drx/config"
	"github.com/krifik/test-drx/constanta"
	"github.com/krifik/test-drx/helper"
	"github.com/krifik/test-drx/model"
	"github.com/krifik/test-drx/module"
	"github.com/krifik/test-drx/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func truncateProduct(db *gorm.DB) {
	db.Raw("TRUNCATE products")
}
func TestProductControllerAddSuccess(t *testing.T) {
	configuration := config.NewConfiguration("../.env.test")
	db := config.NewPostgresDatabase(configuration)
	config.NewRunMigration(db)
	truncateProduct(db)
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	productModule := module.NewProductModule(db)

	routes.Route(app, productModule)
	createProductRequest := model.CreateProductRequest{
		Name:        "Seblak",
		Description: "Seblack yang lebat dan berbiji",
		Quantity:    69,
		Price:       1000,
	}

	requestBody, _ := json.Marshal(createProductRequest)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2000/api/products", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	res, _ := app.Test(request)
	webResp := model.WebResponse{}
	respBody, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(respBody, &webResp)
	assert.Equal(t, 200, webResp.Code)
	assert.Equal(t, "OK", webResp.Status)
	jsonData, _ := json.Marshal(webResp.Data)
	createProductResponse := model.CreateProductResponse{}
	json.Unmarshal(jsonData, &createProductResponse)
	assert.Equal(t, createProductRequest.Name, createProductResponse.Name)
}
func TestProductControllerAddValidationFail(t *testing.T) {
	configuration := config.NewConfiguration("../.env.test")
	db := config.NewPostgresDatabase(configuration)
	truncateProduct(db)
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	productModule := module.NewProductModule(db)

	routes.Route(app, productModule)

	// Invalid request with missing required fields
	invalidProductRequest := model.CreateProductRequest{
		Name:        "",
		Description: "Short",
		Quantity:    0,
		Price:       0,
	}

	requestBody, _ := json.Marshal(invalidProductRequest)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:2000/api/products", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	res, _ := app.Test(request)
	webResp := model.WebResponse{}
	respBody, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(respBody, &webResp)
	assert.Equal(t, fiber.StatusBadRequest, webResp.Code)
	assert.Equal(t, "BAD REQUEST", webResp.Status)
}

func TestApplyDiscount(t *testing.T) {
	discountRequest := model.PriceWithDiscountsRequest{
		OriginalPrice: 250,
		Discounts: []model.Discounts{
			{
				Type:  constanta.DiscountTypeFixed,
				Value: 20,
			},
			{
				Type:  constanta.DiscountTypePercentage,
				Value: 10,
			},
			{
				Type: constanta.DiscountTypeConditional,
				Condition: model.Conditional{
					Amount: 200,
					Value:  15,
				},
			},
			{
				Type: constanta.DiscountTypeTiered,
				Tiers: []model.Tiers{
					{
						Min:   0,
						Max:   99,
						Value: 5,
					},
					{
						Min:   100,
						Max:   199,
						Value: 10,
					},
					{
						Min:   200,
						Max:   9999,
						Value: 25,
					},
				},
			},
			{
				Type:        constanta.DiscountTypeCap,
				MaxDiscount: 60,
			},
		},
	}
	discountRequest2 := model.PriceWithDiscountsRequest{
		OriginalPrice: 300,
		Discounts: []model.Discounts{
			{
				Type:  constanta.DiscountTypeFixed,
				Value: 20,
			},
			{
				Type:  constanta.DiscountTypePercentage,
				Value: 10,
			},
			{
				Type: constanta.DiscountTypeConditional,
				Condition: model.Conditional{
					Amount: 200,
					Value:  15,
				},
			},
			{
				Type: constanta.DiscountTypeTiered,
				Tiers: []model.Tiers{
					{
						Min:   0,
						Max:   99,
						Value: 5,
					},
					{
						Min:   100,
						Max:   199,
						Value: 10,
					},
					{
						Min:   200,
						Max:   9999,
						Value: 25,
					},
				},
			},
			{
				Type:        constanta.DiscountTypeCap,
				MaxDiscount: 60,
			},
		},
	}
	finalPrice := helper.ApplyDiscounts(discountRequest)
	finalPrice2 := helper.ApplyDiscounts(discountRequest2)

	assert.Equal(t, 240.0, finalPrice2)
	assert.Equal(t, 190.0, finalPrice)
}
func TestFixedDiscount(t *testing.T) {
	discountRequest := model.PriceWithDiscountsRequest{
		OriginalPrice: 100,
		Discounts: []model.Discounts{
			{
				Type:  constanta.DiscountTypeFixed,
				Value: 10,
			},
		},
	}
	finalPrice := helper.ApplyDiscounts(discountRequest)
	assert.Equal(t, 90.0, finalPrice)
}

func TestPercentageDiscount(t *testing.T) {
	discountRequest := model.PriceWithDiscountsRequest{
		OriginalPrice: 100,
		Discounts: []model.Discounts{
			{
				Type:  constanta.DiscountTypePercentage,
				Value: 10,
			},
		},
	}
	finalPrice := helper.ApplyDiscounts(discountRequest)
	assert.Equal(t, 90.0, finalPrice)
}

func TestConditionalDiscount(t *testing.T) {
	discountRequest := model.PriceWithDiscountsRequest{
		OriginalPrice: 100,
		Discounts: []model.Discounts{
			{
				Type: constanta.DiscountTypeConditional,
				Condition: model.Conditional{
					Amount: 50,
					Value:  10,
				},
			},
		},
	}
	finalPrice := helper.ApplyDiscounts(discountRequest)
	assert.Equal(t, 90.0, finalPrice)
}

func TestTieredDiscount(t *testing.T) {
	discountRequest := model.PriceWithDiscountsRequest{
		OriginalPrice: 100,
		Discounts: []model.Discounts{
			{
				Type: constanta.DiscountTypeTiered,
				Tiers: []model.Tiers{
					{
						Min:   0,
						Max:   99,
						Value: 5,
					},
					{
						Min:   100,
						Max:   199,
						Value: 10,
					},
					{
						Min:   200,
						Max:   9999,
						Value: 25,
					},
				},
			},
		},
	}
	finalPrice := helper.ApplyDiscounts(discountRequest)
	assert.Equal(t, 90.0, finalPrice)
}

func TestCapDiscount(t *testing.T) {
	discountRequest := model.PriceWithDiscountsRequest{
		OriginalPrice: 100,
		Discounts: []model.Discounts{
			{
				Type:        constanta.DiscountTypeCap,
				MaxDiscount: 5,
			},
		},
	}
	finalPrice := helper.ApplyDiscounts(discountRequest)
	assert.Equal(t, 95.0, finalPrice)
}
