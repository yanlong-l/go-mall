package main

import (
	"errors"
	"github.com/yanlong-l/go-mall/common/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yanlong-l/go-mall/common/logger"
	"github.com/yanlong-l/go-mall/common/middleware"
	"github.com/yanlong-l/go-mall/config"
)

func main() {
	server := gin.Default()
	server.Use(middleware.StartTrace(), middleware.LogAccess(), middleware.GinPanicRecovery())
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
	server.GET("/customized-error-test", func(ctx *gin.Context) {
		// 使用 Wrap 包装原因error 生成 项目error
		err := errors.New("test error-111")
		aAppErr := errcode.Wrap("包装错误", err)
		bAppErr := errcode.Wrap("新包装错误", aAppErr)
		logger.Error(ctx, "打印错误", "err", bAppErr)

		// 预定义的ErrServer, 给其追加错误原因的error
		err = errors.New("a domain error")
		apiErr := errcode.ErrServer.WithCause(err)
		logger.Error(ctx, "API执行中出现错误", "err", apiErr)

		ctx.JSON(apiErr.HttpStatusCode(), gin.H{
			"code": apiErr.Code,
			"msg":  apiErr.Msg,
		})
	})
	server.Run(":9000") // listen and serve on 0.0.0.0:9000
}
