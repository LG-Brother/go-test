package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", defaultHandler)
	r.GET("/testSuccess", successHandler)
	r.GET("/testError", errorHandler)
	// 注意，要用postman、curl等工具测，浏览器默认是用get请求访问
	r.POST("/testPost", postHandler)
	r.Run(":8080")
}

func defaultHandler(c *gin.Context) {
	c.String(http.StatusOK, "This is default page")
}

func successHandler(c *gin.Context) {
	c.String(http.StatusOK, "The request is success!!!")
}

func errorHandler(c *gin.Context) {
	c.String(http.StatusInternalServerError, "The request is error!!!")
}

func postHandler(c *gin.Context) {
	//c.String(http.StatusOK, "Test post request successful!!！")
	c.JSON(http.StatusOK, gin.H{
		"username": "大卫",
		"age":      20,
	})
}
