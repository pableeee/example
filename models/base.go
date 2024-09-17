package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	orm     *gorm.DB
	msqlOrm *gorm.DB
	err     error
)

func InitLite(c db.Connection) {
	orm, err = gorm.Open("sqlite3", c.GetDB("default"))
	if err != nil {
		panic("initialize orm failed")
	}
}

func InitMysql(c db.Connection) {
	msqlOrm, err = gorm.Open("mysql", c.GetDB("mysql"))

	if err != nil {
		panic("initialize orm failed")
	}
}
