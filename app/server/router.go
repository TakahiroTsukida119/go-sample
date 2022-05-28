package server

import (
	"github.com/TakahiroTsukida119/go-sample.git/internal/interface/handler"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(e *gin.Engine) *gin.Engine {
	v1 := e.Group("/api/v1")
	{
		userGroup := v1.Group("/user")
		{
			userGroup.GET("/", handler.Index)
		}
	}
	return e
}
