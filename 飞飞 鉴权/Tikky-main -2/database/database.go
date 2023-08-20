package database

import (
	"Tikky-main-2/controller"
	"fmt"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "你的数据库名"
	username := "你的mysql用户名"
	password := "你的mysql密码"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

	//迁移
	db.AutoMigrate(&controller.User{})

	return db

}
