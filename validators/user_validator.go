package validators

import (
	"Demo/models"

	"github.com/go-playground/validator/v10"
)

type IUserValidator interface {
	Validate(userAuth models.UserAuth) error
}

type userValidator struct {
	v *validator.Validate
}

func InitUserValidator(v *validator.Validate) IUserValidator {
	return &userValidator{v}
}

func (uv *userValidator) Validate(userAuth models.UserAuth) error {
	return uv.v.Struct(userAuth)
}
