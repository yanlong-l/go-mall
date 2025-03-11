package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yanlong-l/go-mall/common/logger"
	"github.com/yanlong-l/go-mall/config"
)

func main() {
	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		logger.ZapLoggerTest(ctx)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Pong",
		})
	})
	server.GET("config-read", func(ctx *gin.Context) {
		db := config.Database
		app := config.App
		ctx.JSON(http.StatusOK, gin.H{
			"dsn":         db.DSN,
			"maxlifetime": db.MaxLifeTime,
			"app_name":    app.Name,
			"app_env":     app.Env,
		})
	})
	server.Run(":9000")
}
