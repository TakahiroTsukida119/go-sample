package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Sample(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "sample!!"})
}
