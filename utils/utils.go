package utils

import (
	"github.com/anisbhsl/fetch-backend-assignment/models"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func RegisterValidator() {
	validate = validator.New(validator.WithRequiredStructEnabled())

	// register all custom validators
	validate.RegisterValidation("validateRetailer", models.ValidateRetailer)
	validate.RegisterValidation("validateReceiptTotal", models.ValidateReceiptTotal)
	validate.RegisterValidation("validateReceiptItemShortDesc", models.ValidateReceiptItemShortDesc)
	validate.RegisterValidation("validateReceiptItemPrice", models.ValidateReceiptItemPrice)
}

func GetValidator() *validator.Validate {
	return validate
}
