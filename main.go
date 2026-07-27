package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "You're Server is up and running!",
		})
	})
	server.Run(":8080")

}
