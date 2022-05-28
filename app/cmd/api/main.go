package main

import (
	"github.com/TakahiroTsukida119/go-sample.git/config"
	"github.com/TakahiroTsukida119/go-sample.git/middleware"
	"github.com/TakahiroTsukida119/go-sample.git/server"
	"github.com/gin-gonic/gin"
)

var ginLocal *gin.Engine

func init() {
	ginLocal = NewGin()
}

func main() {
	config.DBOpen()
	defer config.DBClose()
	if err := ginLocal.Run(":8000"); err != nil {
		return
	}
}

func NewGin() *gin.Engine {
	e := gin.New()
	e.Use(middleware.Cors())
	e = server.ApiRoutes(e)
	return e
}
