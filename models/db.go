package models

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type connection struct{}

func New() *connection {
	return &connection{}
}

// https://gorm.io/docs/connecting_to_the_database.html
func (c *connection) InitDB() *gorm.DB {
	dsn := viper.GetString(`mysql.dsn`)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to initialize DB!!!")
		log.Fatal(err)
	}
	return db
}
