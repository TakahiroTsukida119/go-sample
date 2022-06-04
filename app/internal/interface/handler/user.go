package handler

import (
	"github.com/TakahiroTsukida119/go-sample.git/model"
	"github.com/TakahiroTsukida119/go-sample.git/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler interface {
	Index(ctx *gin.Context)
}

type userHandler struct {
	us service.UserService
}

func NewUserHandler(us service.UserService) UserHandler {
	return &userHandler{
		us: us,
	}
}

type ListUserResponse struct {
	Users []*model.User `json:"users"`
}

// Index godoc
// @router /user [get]
// @Security AuthToken
// @Summary ユーザー一覧
// @description ユーザーの一覧を取得します
// @accept application/json
// @Param email query string false "メールアドレス"
// @Success 200 {object} ListUser
// @Failure 400 {object} helper.Error
func (uh *userHandler) Index(ctx *gin.Context) {
	users, err := uh.us.Index()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, &ListUserResponse{
		Users: users,
	})
}
