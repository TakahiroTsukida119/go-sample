package service

import (
	"github.com/TakahiroTsukida119/go-sample.git/model"
	"github.com/TakahiroTsukida119/go-sample.git/repository"
)

type UserService interface {
	Index() ([]*model.User, error)
}

type userService struct {
	Name string
	ur   repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		Name: "UserService",
		ur:   ur,
	}
}

func (us *userService) Index() ([]*model.User, error) {
	_, err := us.ur.FetchAllUsers()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
