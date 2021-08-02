package internal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/marckii8888/TAP_meteor/config"
)

var (
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME  string
	DB_HOST  string
	DB_PORT  string
)

var db *gorm.DB

func InitDB() *gorm.DB {
	DB_USERNAME = config.Conf.Database.User
	DB_PASSWORD = config.Conf.Database.Password
	DB_NAME = config.Conf.Database.Name
	DB_HOST = config.Conf.Database.Host
	DB_PORT = config.Conf.Database.Port
	db = connectDB()
	return db
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database : error=%v\n", err)
		return nil
	}
	return db
}