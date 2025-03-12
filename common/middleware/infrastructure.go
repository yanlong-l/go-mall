package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yanlong-l/go-mall/common/util"
)

// infrastructure 中存放项目运行需要的基础中间件

func StartTrace() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.Request.Header.Get("traceid")
		pSpanId := c.Request.Header.Get("spanid")
		spanId := util.GenerateSpanID(c.Request.RemoteAddr)
		if traceId == "" { // 如果traceId 为空，证明是链路的发端，把它设置成此次的spanId，发端的spanId是root spanId
			traceId = spanId // trace 标识整个请求的链路, span则标识链路中的不同服务
		}
		c.Set("traceid", traceId)
		c.Set("spanid", spanId)
		c.Set("pspanid", pSpanId)
		c.Next()
	}
}
