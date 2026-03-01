package router

import (
	"fmt"
	"gospaacex/bff/handler/service"
	"time"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	pos := r.Group("/pos")
	{
		pos.POST("/create", service.PosCreate)
		pos.POST("/del", service.PosDel)
		pos.POST("/update", service.PosUpdate)
		pos.GET("/list", service.PosList)
	}
	return r
}
