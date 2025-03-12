package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yanlong-l/go-mall/common/logger"
	"github.com/yanlong-l/go-mall/common/middleware"
	"github.com/yanlong-l/go-mall/config"
)

func main() {
	server := gin.Default()
	server.Use(gin.Logger(), middleware.StartTrace())
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Pong",
		})
	})
	server.GET("/config-read", func(ctx *gin.Context) {
		db := config.Database
		app := config.App
		ctx.JSON(http.StatusOK, gin.H{
			"dsn":         db.DSN,
			"maxlifetime": db.MaxLifeTime,
			"app_name":    app.Name,
			"app_env":     app.Env,
		})
	})
	server.GET("/logger-test", func(ctx *gin.Context) {
		logger.Info(ctx, "logger test", "name", "zhangsan", "age", 18)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	server.Run(":9000") // listen and serve on 0.0.0.0:9000
}
