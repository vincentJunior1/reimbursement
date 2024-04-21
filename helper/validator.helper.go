package helper

import (
	"github.com/go-playground/validator/v10"
)

func Validatestruct(payload any) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(payload)

	return err
}
