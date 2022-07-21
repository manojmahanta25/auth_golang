package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	ID        int64      `gorm:"primary_key;auto_increment;not_null" json:"id"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime;default:now()"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime;default:now()"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
}

var DBConn *gorm.DB

func Conn() {
	//open a db connection
	var err error
	DBConn, err = gorm.Open("mysql", Config("DB_USER")+":"+Config("DB_PASS")+"@tcp("+Config("DB_HOST")+":"+Config("DB_PORT")+")/"+Config("DB_NAME")+"?charset=utf8&parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema

}
