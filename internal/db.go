package internal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "root123"
const DB_NAME = "meteor_assessment"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var db *gorm.DB

func InitDB() *gorm.DB {
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