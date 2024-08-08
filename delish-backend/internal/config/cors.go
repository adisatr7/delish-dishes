package config

import (
	"time"

	"github.com/gin-contrib/cors"
)

var CORS = cors.Config{
	AllowOrigins:     []string{"http://localhost:4000"},
	AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
	AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
	MaxAge:           12 * time.Hour,
}
