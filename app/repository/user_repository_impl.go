package repository

import (
	"github.com/TakahiroTsukida119/go-sample.git/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FetchAllUsers() (*[]model.User, error)
}

type userRepository struct {
	client *gorm.DB
}

func NewUserRepository(client *gorm.DB) UserRepository {
	return &userRepository{client: client}
}

func (ur *userRepository) FetchAllUsers() (*[]model.User, error) {
	var _ []*model.User
	//if err := ur.client.DB().(&users); err != nil {
	//
	//}

	return nil, nil
}
