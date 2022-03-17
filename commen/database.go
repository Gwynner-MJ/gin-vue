package commen

import (
	"fmt"
	"gin-vue/model"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginvue"
	username := "root"
	password := "1120MaJunNan."
	charset := "utf8"
	arg := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, arg)
	if err != nil {
		panic("打开数据库失败,err" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
