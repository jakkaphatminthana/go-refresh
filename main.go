package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakkaphatminthana/go-refresh/config"
	"github.com/jakkaphatminthana/go-refresh/middlewares"
	"github.com/jakkaphatminthana/go-refresh/utils"
)

func InitLoadEnv() {
	// Load config from env
	appConfig, errLoadConfig := config.LoadConfig("./config/env")
	if errLoadConfig != nil {
		log.Fatalf("Fatal error: could not load configuration. %v", errLoadConfig)
	}
	config.AppConfig = appConfig
}

func main() {
	app := gin.Default()

	app.Use(middlewares.CORS())

	InitLoadEnv()
	utils.InitializeLogger()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"value":   config.AppConfig.DatabaseName,
		})
	})

	app.Run(":8080")
}
