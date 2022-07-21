package userInterface

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type LoginBody struct {
	Email string `json:"email" xml:"email" form:"email"`
	Pass  string `json:"pass" xml:"pass" form:"pass"`
}

func (p LoginBody) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Email, validation.Required, validation.Length(4, 255)),
		validation.Field(&p.Pass, validation.Required, validation.Length(8, 50)),
	)
}
