package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func main() {
	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Hello, World!",
		})
	})

	app.Run()
}