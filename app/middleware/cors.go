package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:8080",
		},
		AllowHeaders: []string{
			"Accept",
			"Accept-Encoding",
			"Authorization",
			"Content-Type",
			"Content-Length",
			"Origin",
			"Referer",
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		MaxAge: 24 * time.Hour,
	})
}
