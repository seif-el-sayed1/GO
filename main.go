package main

import (
	"github.com/gin-gonic/gin"
	"github.com/seif-el-sayed1/GO/db"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "You're Server is up and running!",
		})
	})
	server.Run(":8080")

}
