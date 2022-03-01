package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type connection struct{}

func New() *connection {
	return &connection{}
}

// https://gorm.io/docs/connecting_to_the_database.html
func (c *connection) InitDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/mini_chat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to initialize DB!!!")
		log.Fatal(err)
	}
	return db
}
