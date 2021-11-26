package main

import (
	"goact/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		core.LoadEnv()
		core.NewDatabase()
		context.JSON(http.StatusOK, gin.H{"data": "Hello World !"})
	})
	router.Run(":8000")
}
