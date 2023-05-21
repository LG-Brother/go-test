package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "This is Home Page")
	})
	r.GET("/index", func(context *gin.Context) {
		context.String(http.StatusOK, "This is index page")
	})
	r.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello %s", name)
	})
	r.GET("/users", func(context *gin.Context) {
		name := context.Query("name")
		role := context.DefaultQuery("role", "teacher")
		context.String(http.StatusOK, "%s is a %s", name, role)
	})
	// 测试命令：curl http://localhost:8080/form -X POST -d 'username=林刚&password=33se2'
	r.POST("/form", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.DefaultPostForm("password", "123456")
		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	// 测试命令：curl -g "http://localhost:8080/post?ids[tom]=001&ids[lucy]=002" -X POST -d 'names[a]=sam&names[b]=David'
	// -g表示允许使用特殊字符
	r.POST("/post", func(context *gin.Context) {
		ids := context.QueryMap("ids")
		names := context.PostFormMap("names")
		context.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	// 访问路径会变的重定向
	r.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/index")
	})
	// 访问路径不变的重定向
	r.GET("/goIndex", func(context *gin.Context) {
		context.Request.URL.Path = "/"
		r.HandleContext(context)
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/one", defaultHandler1)
		v1.GET("/two", defaultHandler1)
	}
	r.Run(":8080")
}

func defaultHandler1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"path": c.FullPath(),
	})
}
