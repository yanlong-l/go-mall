package main

import (
	"github.com/yanlong-l/go-mall/api/router"
	"github.com/yanlong-l/go-mall/common/enum"
	"github.com/yanlong-l/go-mall/config"

	"github.com/gin-gonic/gin"
)

func main() {
	if config.App.Env == enum.ModeProd {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	router.RegisterRoutes(server)
	server.Run(":9000") // listen and serve on 0.0.0.0:9000
}
