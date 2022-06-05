package main

import (
	"github.com/TakahiroTsukida119/go-sample.git/config"
	"github.com/TakahiroTsukida119/go-sample.git/internal/interface/handler"
	"github.com/TakahiroTsukida119/go-sample.git/middleware"
	"github.com/TakahiroTsukida119/go-sample.git/repository"
	"github.com/TakahiroTsukida119/go-sample.git/server"
	"github.com/TakahiroTsukida119/go-sample.git/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ginLocal *gin.Engine

func main() {
	sqlHandler := config.DBOpen()
	defer config.DBClose()
	ginLocal = NewGin(sqlHandler.DB)
	if err := ginLocal.Run(":8000"); err != nil {
		return
	}
}

func NewGin(client *gorm.DB) *gin.Engine {
	//ur := provider.DIContainer.MustGet("userRepository").(ri.UserRepository)
	//us := provider.DIContainer.MustGet("userService").(si.UserService)
	// repository
	ur := repository.NewUserRepository(client)
	// service
	us := service.NewUserService(ur)
	// handler
	uh := handler.NewUserHandler(us)
	// gin
	e := gin.New()
	e.Use(middleware.Cors())
	e = server.ApiRoutes(e, uh)
	return e
}
