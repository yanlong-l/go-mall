package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yanlong-l/go-mall/common/app"
	"github.com/yanlong-l/go-mall/common/errcode"
	"github.com/yanlong-l/go-mall/common/logger"
	"github.com/yanlong-l/go-mall/config"
)

func TestPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func TestConfigRead(ctx *gin.Context) {
	db := config.Database
	app := config.App
	ctx.JSON(http.StatusOK, gin.H{
		"dsn":         db.Master.DSN,
		"maxlifetime": db.Master.MaxLifeTime,
		"app_name":    app.Name,
		"app_env":     app.Env,
	})
}

func TestLogger(ctx *gin.Context) {
	logger.Info(ctx, "logger test", "name", "zhangsan", "age", 18)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func TestPanicLog(c *gin.Context) {
	var a map[string]string
	a["k"] = "v"
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   a,
	})
}

func TestAppError(ctx *gin.Context) {
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
}

func TestResponseObj(c *gin.Context) {
	data := map[string]int{
		"a": 1,
		"b": 2,
	}
	app.NewResponse(c).Success(data)
}

func TestResponseList(c *gin.Context) {
	pagination := app.NewPagination(c)
	// Mock fetch list data from db
	data := []struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		{
			Name: "Lily",
			Age:  26,
		},
		{
			Name: "Violet",
			Age:  25,
		},
	}
	pagination.SetTotalRows(2)
	app.NewResponse(c).SetPagination(pagination).Success(data)
}

func TestResponseError(c *gin.Context) {
	baseErr := errors.New("a dao error")
	// 这一步正式开发时写在service层
	err := errcode.Wrap("encountered an error when xxx service did xxx", baseErr)
	app.NewResponse(c).Error(errcode.ErrServer.WithCause(err))
}

func TestAccessLog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
