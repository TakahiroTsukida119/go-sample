package repository

import (
	"github.com/TakahiroTsukida119/go-sample.git/model"
	"github.com/TakahiroTsukida119/go-sample.git/repository/interfaces"
	"gorm.io/gorm"
)

type UserRepository struct {
	client *gorm.DB
}

func NewUserRepository(client *gorm.DB) interfaces.UserRepository {
	return &UserRepository{client: client}
}

func (r UserRepository) Index() ([]model.User, error) {
	var users []model.User

	if err := r.client.First(&users); err != nil {

	}

	return users, nil
}
