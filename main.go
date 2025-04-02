package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yanlong-l/go-mall/api/router"
	"github.com/yanlong-l/go-mall/common/enum"
	"github.com/yanlong-l/go-mall/config"
	_ "github.com/yanlong-l/go-mall/dal/cache"
	_ "github.com/yanlong-l/go-mall/dal/dao"
)

func main() {
	if config.App.Env == enum.ModeProd {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	router.RegisterRoutes(server)
	server.Run(":9000") // listen and serve on 0.0.0.0:9000
}
