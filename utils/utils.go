package utils

import (
	"sync"

	"github.com/anisbhsl/fetch-backend-assignment/models"
	"github.com/go-playground/validator/v10"
)

var (
	validate     *validator.Validate
	validateOnce sync.Once
)

// registerValidator registers all the custom validators
func registerValidator() {
	validate = validator.New(validator.WithRequiredStructEnabled())

	// register all custom validators
	validate.RegisterValidation("validateRetailer", models.ValidateRetailer)
	validate.RegisterValidation("validateReceiptTotal", models.ValidateReceiptTotal)
	validate.RegisterValidation("validateReceiptItemShortDesc", models.ValidateReceiptItemShortDesc)
	validate.RegisterValidation("validateReceiptItemPrice", models.ValidateReceiptItemPrice)
}

// GetValidator returns an instance of a validator
func GetValidator() *validator.Validate {
	validateOnce.Do(registerValidator)
	return validate
}
