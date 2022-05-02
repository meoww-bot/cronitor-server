package main

import (
	handler "cronitor-server/handler"
	"io"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
)

func setupLogOutput() {

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	router.GET("/", handler.Root)

	router.GET("/ping/:apiKey/:code", handler.Ping) // cronitor exec

	api := router.Group("/api") // for web ui
	{
		api.GET("/monitors", handler.GetAllMonitors)
		api.GET("/monitors/:monitor_code", handler.GetMonitorInfo)
		api.GET("/aggregates", handler.GetMonitorAggregates)

	}

	v3 := router.Group("/v3") // for cronitor-cli
	{
		v3.GET("/", handler.Root)
		v3.PUT("/monitors", handler.PutMonitorsV3)
		v3.GET("/monitors", handler.GetMonitorsV3)
		v3.GET("/monitors/:monitor_code/activity", handler.GetMonitorActivityV3)
		v3.GET("/user/:username", handler.GetUserInfo)
		v3.POST("/user/:username", handler.CreateUser)
	}

	// 如果应用程序不在代理之后，“ForwardedByClientIP”应设置为 false，因此“X-Forwarded-For”将被忽略。
	// 如果在代理后面将其设置为true
	router.ForwardedByClientIP = false

	router.Run(":8000")
}
