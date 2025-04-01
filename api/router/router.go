package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yanlong-l/go-mall/common/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	// use global middlewares
	server.Use(middleware.StartTrace(), middleware.LogAccess(), middleware.GinPanicRecovery())
	routerGroup := server.Group("")
	registerBuildingRoutes(routerGroup)
}
