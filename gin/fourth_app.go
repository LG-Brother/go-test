package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type student struct {
	Name string
	Age  int
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/arr", arrHandler)
	r.Run(":8080")
}

func arrHandler(context *gin.Context) {
	stu1 := &student{Name: "Tom", Age: 18}
	stu2 := &student{Name: "Lucy", Age: 28}
	context.HTML(http.StatusOK, "arr.html", gin.H{
		"title":  "Gin",
		"stuArr": [2]*student{stu1, stu2},
	})
}
