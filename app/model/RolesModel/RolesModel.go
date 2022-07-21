package RolesModel

import (
	"authMicroservice/app/config"
	roleInterface2 "authMicroservice/app/http/interfaces/roleInterface"
	"authMicroservice/app/model/UserPermissionModel"
	"errors"
	"time"
)

type Role struct {
	config.Model
	Name       string                              `json:"name" gorm:"not null"`
	Permission UserPermissionModel.PermissionArray `json:"permissions,omitempty" gorm:"type:json;not null;"`
}

func Migrate() {
	config.DBConn.AutoMigrate(&Role{})
}

func FindById(roleId int) (Role, error) {
	var db = config.DBConn
	var role Role
	if err := db.Find(&role, roleId).Error; err != nil {
		err := errors.New("user not found")
		return role, err
	}
	return role, nil
}

func AddRole(roleInterface roleInterface2.RoleInterface) (Role, error) {
	var db = config.DBConn
	var role Role
	role.Name = roleInterface.Name
	role.Permission = roleInterface.Permission
	if result := db.Create(&role); result.Error != nil {
		return role, result.Error
	}
	return role, nil
}

func UpdateRole(roleInterface roleInterface2.RoleInterface, roleId int) (Role, error) {
	var db = config.DBConn
	var role Role
	role, err := FindById(roleId)
	if err != nil {
		return role, err
	}
	role.Name = roleInterface.Name
	role.Permission = roleInterface.Permission
	role.UpdatedAt = time.Now()
	if result := db.Save(&role); result.Error != nil {
		return role, result.Error
	}
	return role, nil
}

func DeleteRole(roleId int) (Role, error) {
	var db = config.DBConn
	var role Role
	role, err := FindById(roleId)
	if err != nil {
		return role, err
	}
	if role.Name == "SuperAdmin" {
		err = errors.New("not authorized to delete this role")
		return role, err
	}
	if result := db.Delete(&role); result.Error != nil {
		return role, result.Error
	}
	return role, nil
}
