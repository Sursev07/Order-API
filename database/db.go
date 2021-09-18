package database

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "root"
	password = ""
	dbPort   = "3306"
	dbname   = "orders_by"
	db       *gorm.DB
	err      error
)

func InitDB() {
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",user, password, host, dbPort, dbname)
	dsn := config
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("====successfully connected to database=====")

}

func GetDB() *gorm.DB {
	return db
}