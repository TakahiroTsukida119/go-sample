package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index godoc
// @router /user [get]
// @Security AuthToken
// @Summary ユーザー一覧
// @description ユーザーの一覧を取得します
// @accept application/json
// @Param email query string false "メールアドレス"
// @Success 200 {object} ListUser
// @Failure 400 {object} helper.Error
func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "user index!!"})
}
