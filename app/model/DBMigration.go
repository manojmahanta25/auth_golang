package model

import (
	"authMicroservice/app/model/RolesModel"
	"authMicroservice/app/model/UserModel"
	"authMicroservice/app/model/UserPermissionModel"
)

func Migrate() {
	UserModel.Migrate()
	RolesModel.Migrate()
	UserPermissionModel.Migrate()
}
