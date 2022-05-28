package interfaces

import "github.com/TakahiroTsukida119/go-sample.git/model"

type UserService interface {
	Index() (model.User, error)
}
