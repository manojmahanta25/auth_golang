package userInterface

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"regexp"
)

type SignUpBody struct {
	Email  string `json:"email" xml:"email" form:"email"`
	Pass   string `json:"password" xml:"password" form:"password"`
	FName  string `json:"first_name" xml:"first_name" form:"first_name"`
	LName  string `json:"last_name" xml:"last_name" form:"last_name"`
	Gender string `json:"gender" xml:"gender" form:"gender"`
	Phone  string `json:"phone" xml:"phone" form:"phone"`
	DOB    string `json:"dob" xml:"dob" form:"dob"`
}

func (p SignUpBody) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.FName, validation.Required, validation.Length(2, 20), validation.Match(regexp.MustCompile(`^[a-zA-Z \-\s]+$`))),
		validation.Field(&p.LName, validation.Required, validation.Length(2, 20), validation.Match(regexp.MustCompile(`^[a-zA-Z \-\s]+$`))),
		validation.Field(&p.Email, validation.Required, validation.Length(4, 255), is.Email),
		validation.Field(&p.Pass, validation.Required, validation.Length(8, 50)),
		validation.Field(&p.Gender, validation.Required, validation.Match(regexp.MustCompile(`^([0-3]{1})+$`))),
		validation.Field(&p.Phone, validation.When(p.Phone != "", validation.Match(regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)).Error("Invalid Phone Number"))),
		validation.Field(&p.DOB, validation.Required, validation.Match(regexp.MustCompile(`^([0-2][0-9]|(3)[0-1])(\-)(((0)[0-9])|((1)[0-2]))(\-)\d{4}$`))),
	)
}
