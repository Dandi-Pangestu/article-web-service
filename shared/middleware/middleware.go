package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	AllowedOrigins = []string{"*"}
	AllowedMethods = []string{"POST", "OPTIONS", "GET", "PUT", "PATCH", "DELETE"}
	AllowedHeaders = []string{
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
		"Accept",
		"Origin",
		"Cache-Control",
		"X-Requested-With",
		"Client-Key",
	}
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     AllowedOrigins,
		AllowMethods:     AllowedMethods,
		AllowHeaders:     AllowedHeaders,
		AllowCredentials: true,
	})
}
