package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakkaphatminthana/go-refresh/config"
	"github.com/jakkaphatminthana/go-refresh/database"
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

func init() {
	InitLoadEnv()
	utils.InitializeLogger()

	// database
	if _, errConnectDB := database.ConnectDB(); errConnectDB != nil {
		panic(errConnectDB)
	} else {
		fmt.Println("[INFO] 📦 Initialize Connect database successfully...")
	}
}

func main() {
	app := gin.Default()

	// For security
	if err := app.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	app.Use(middlewares.CORS())

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"value":   config.AppConfig.DatabaseName,
		})
	})

	app.Run(":8080")
}
