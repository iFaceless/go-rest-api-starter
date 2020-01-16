package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

var (
	allowedOrigins = []string{
		"localhost", // just an example
	}
)

func DefaultCORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = false
	config.AllowCredentials = true
	config.AllowOriginFunc = func(origin string) bool {
		for _, o := range allowedOrigins {
			if strings.Contains(origin, o) {
				return true
			}
		}
		return false
	}
	return cors.New(config)
}
