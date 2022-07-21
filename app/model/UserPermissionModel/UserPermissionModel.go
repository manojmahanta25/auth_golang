package UserPermissionModel

import (
	"authMicroservice/app/config"
	"database/sql/driver"
	"encoding/json"
)

type ObjectPermission struct {
	Name       string     `json:"name"`
	Permission Permission `json:"permission"`
}
type Permission struct {
	Read   bool `json:"read"`
	Write  bool `json:"write"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}
type PermissionArray []ObjectPermission
type UserPermission struct {
	config.Model
	UserId      int64           `json:"user_id" gorm:"index;unique"`
	Permissions PermissionArray `json:"permissions" gorm:"type:json;not null;"`
}

func Migrate() {
	config.DBConn.AutoMigrate(&UserPermission{})
}
func (m *PermissionArray) Scan(src interface{}) error {
	val := src.([]byte)
	return json.Unmarshal(val, &m)
}
func (m PermissionArray) Value() (driver.Value, error) {
	j, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return driver.Value([]byte(j)), nil
}
func FindByUserId(userId int) (UserPermission, error) {
	var userPermission UserPermission
	var db = config.DBConn
	err := db.Find(&userPermission, userId).Error
	if err != nil {
		return userPermission, err
	}
	return userPermission, nil
}
