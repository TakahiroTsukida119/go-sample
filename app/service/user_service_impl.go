package service

import (
	"github.com/TakahiroTsukida119/go-sample.git/model"
	"github.com/TakahiroTsukida119/go-sample.git/service/interfaces"
)

type UserService struct {
}

func NewUserService() interfaces.UserService {
	return &UserService{}
}

func (u UserService) Index() (model.User, error) {
	return model.User{
		Id:        "id",
		Name:      "sample",
		Email:     "example@test.test",
		CreatedAt: "",
		UpdatedAt: "",
		DeletedAt: "",
	}, nil
}
