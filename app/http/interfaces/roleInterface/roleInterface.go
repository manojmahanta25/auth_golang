package roleInterface

import (
	"authMicroservice/app/model/UserPermissionModel"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RoleInterface struct {
	Name       string                              `json:"name" xml:"name" form:"name"`
	Permission UserPermissionModel.PermissionArray `json:"permissions" xml:"permissions" form:"permissions"`
}

func (p RoleInterface) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&p.Permission, validation.Required),
	)
}
