package main

import (
	"html/template"
	"io"
	"mini-chat/models"
	"mini-chat/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigType("json")

	// init log
	gin.DisableConsoleColor()
	logfile := "/logs/gin.log"
	f, _ := os.Create(logfile)
	gin.DefaultWriter = io.MultiWriter(f)

	// init DB
	models.New().InitDB()
}

func StartRoute() {
	router := routers.Router()
	html := template.Must(template.ParseFiles("/view/index.html"))
	router.SetHTMLTemplate(html)
}

func main() {
	go Init()
	go StartRoute()
}
