package main

import (
	"html/template"
	"io"
	"mini-chat/models"
	"mini-chat/routers"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigType("json")

	// init log
	gin.DisableConsoleColor()
	timeNow := time.Now().Format("2006-01-02_15_04_05")
	logfile := "logs/gin_" + timeNow + ".log"
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
	Init()
	StartRoute()
}
