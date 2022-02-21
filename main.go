package main

import (
	"mini-chat/routers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("json")
}

func main() {
	router := gin.Default()
	routers.Router()
}
