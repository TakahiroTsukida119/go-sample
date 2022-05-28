package interfaces

import "github.com/TakahiroTsukida119/go-sample.git/model"

type UserRepository interface {
	Index() ([]model.User, error)
}
