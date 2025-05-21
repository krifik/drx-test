package validation

import (
	"github.com/krifik/test-drx/exception"
	"github.com/krifik/test-drx/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

func Validate(request model.CreateProductRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Price, validation.Required, validation.Min(1)),
		validation.Field(&request.Quantity, validation.Required, validation.Min(1)),
		validation.Field(&request.Description, validation.Required, validation.Length(10, 10000000)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
