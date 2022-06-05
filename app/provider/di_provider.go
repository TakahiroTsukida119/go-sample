package provider

import (
	"github.com/TakahiroTsukida119/go-sample.git/internal/infrastructure/db"
	"github.com/TakahiroTsukida119/go-sample.git/repository"
	"github.com/TakahiroTsukida119/go-sample.git/service"
	"github.com/fgrosse/goldi"
)

func registerTypes(types goldi.TypeRegistry) {
	types.RegisterAll(map[string]goldi.TypeFactory{
		"newDBClient": goldi.NewType(db.GetNewDBClient),
		// repository
		"userRepository": goldi.NewType(repository.NewUserRepository, "@newDBClient"),
		// service
		"userService": goldi.NewType(service.NewUserService, "@userRepository"),
	})
}
