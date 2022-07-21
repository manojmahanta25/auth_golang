package roleController

import (
	roleInterface2 "authMicroservice/app/http/interfaces/roleInterface"
	"authMicroservice/app/model/RolesModel"
	"authMicroservice/app/utils/handlers"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func CreateRole(c *fiber.Ctx) error {
	var roleInterface roleInterface2.RoleInterface
	if err := c.BodyParser(&roleInterface); err != nil {
		fmt.Println("Error", err.Error())
		return handlers.ErrorJsonOutput(c, err, 400)
	}
	if err := roleInterface.Validate(); err != nil {
		return handlers.ErrorJsonOutput(c, err, 400)
	}
	roleAdd, err := RolesModel.AddRole(roleInterface)
	return handlers.JsonOutputOrError(c, roleAdd, err, 201, 500)
}

func UpdateRole(c *fiber.Ctx) error {
	//var roleInterface roleInterface2.RoleInterface
	return handlers.JsonOutputOrError(c, "roleAdd", nil, 201, 500)
}
