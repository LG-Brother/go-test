package main

import (
	"github.com/gin-gonic/gin"
)

func helloHandler(context *gin.Context) {
	context.String(200, "This is my first Gin application")
}
func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(200, "Hello, Gin")
	})
	r.GET("/hello", helloHandler)
	r.GET("/error", func(context *gin.Context) {
		context.String(500, "This is the error page")
	})
	r.Run(":8081")
}
