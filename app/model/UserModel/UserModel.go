package UserModel

import (
	"authMicroservice/app/config"
	"authMicroservice/app/http/interfaces/userInterface"
	"authMicroservice/app/model/RolesModel"
	"authMicroservice/app/model/UserPermissionModel"
	"authMicroservice/app/utils/helpers"
	"errors"
	"strings"
	"time"
)

type User struct {
	config.Model
	OAuthProvider string                              `json:"o_auth_provider" gorm:"not null"`
	OAuthId       *int                                `json:"o_auth_id,omitempty" gorm:"type:bigint(20);default:null"`
	GivenName     string                              `json:"given_name" gorm:"not null"`
	FamilyName    string                              `json:"family_name" gorm:"not null"`
	Email         string                              `json:"email" gorm:"unique;not null"`
	DateOfBirth   time.Time                           `json:"date_of_birth" gorm:"type:date"`
	IsEmailValid  *bool                               `json:"is_email_valid" gorm:"default:false;"`
	Phone         string                              `json:"phone" gorm:"type:varchar(15);unique;default:null"`
	IsPhoneValid  *bool                               `json:"is_phone_valid" gorm:"default:false;"`
	Gender        string                              `json:"gender" gorm:"type:varchar(1);default:0"`
	Picture       *string                             `json:"picture,omitempty"`
	Password      *string                             `json:"-"`
	LastIp        *string                             `json:"last_ip,omitempty"`
	Roles         []*RolesModel.Role                  `json:"roles,omitempty" gorm:"many2many:user_roles;foreignKey:ID;joinForeignKey:UserRolesId;References:ID;joinReferences:RolesId"`
	Permission    *UserPermissionModel.UserPermission `json:"permissions,omitempty" gorm:"foreignKey:UserId;"`
}

func Migrate() {
	config.DBConn.AutoMigrate(&User{})
}

func GetGender(gender string) string {
	gender = strings.Trim(gender, " ")
	switch gender {
	case "0":
		return "Male"
	case "1":
		return "Female"
	case "2":
		return "Others"
	default:
		return "Disclosed"
	}
}

func FindById(id int) (User, error) {
	var db = config.DBConn.Model(&User{})
	var user User
	if err := db.Preload("Roles").Preload("Permission").Find(&user, id).Error; err != nil {
		err := errors.New("user not found")
		return user, err
	}
	return user, nil
}
func UpdateLastIP(id int, ip string) bool {
	var db = config.DBConn
	if err := db.Model(&User{}).Where("id = ?", id).Update("last_ip", ip).Error; err != nil {
		return false
	}
	return true
}
func LoginByEmail(email string, pass string, ip string) (User, error) {
	var db = config.DBConn
	var user User
	if err := db.Where("email=?", email).First(&user).Error; err != nil {
		err := errors.New("invalid credentials")
		return user, err
	}
	bool := helpers.CheckPasswordHash(pass, *user.Password)
	if bool {
		UpdateLastIP(int(user.ID), ip)
		return user, nil
	} else {
		err := errors.New("invalid password")
		return user, err
	}
}
func LoginByPhone(phone int, pass string, ip string) (User, error) {
	var db = config.DBConn
	var user User
	if err := db.Where("phone=?", phone).First(&user).Error; err != nil {
		err = errors.New("invalid credentials")
		return user, err
	}
	bool := helpers.CheckPasswordHash(pass, *user.Password)
	if bool {
		UpdateLastIP(int(user.ID), ip)
		return user, nil
	} else {
		err := errors.New("invalid password")
		return user, err
	}
}
func SignUpByLocal(signUp userInterface.SignUpBody) (User, error) {
	var db = config.DBConn
	var user User
	password, err := helpers.HashPassword(signUp.Pass)
	if err != nil {
		return user, err
	}
	user.OAuthProvider = "local"
	user.GivenName = signUp.FName
	user.FamilyName = signUp.LName
	user.Email = signUp.Email
	user.Phone = signUp.Phone
	user.Gender = signUp.Gender
	user.Password = &password
	myDate, err := time.Parse("02-01-2006", signUp.DOB)
	if err != nil {
		return user, err
	}
	user.DateOfBirth = myDate
	if result := db.Create(&user); result.Error != nil {
		return user, result.Error
	}
	return user, nil

}
func FindUserByEmail(email string) (User, error) {
	var db = config.DBConn
	var user User
	if err := db.Where("email=?", email).First(&user).Error; err != nil {
		err = errors.New("user not found")
		return user, err
	}
	return user, nil
}
