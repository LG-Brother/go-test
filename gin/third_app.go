package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello %s", name)
	})
	r.GET("/users", func(context *gin.Context) {
		name := context.Query("name")
		role := context.DefaultQuery("role", "teacher")
		context.String(http.StatusOK, "%s is a %s", name, role)
	})
	r.POST("/form", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.DefaultPostForm("password", "123456")
		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	r.Run(":8080")
}
