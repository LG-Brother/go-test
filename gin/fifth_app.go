package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	// 作用于全局
	r.Use(Logger1())
	r.GET("/test-quan", func(context *gin.Context) {
		println("作用于全局路由....")
	})
	// 作用于单路由
	r.GET("/test-dan", Logger2(), func(context *gin.Context) {
		println("作用于单个路由....")
	})
	// 作用于路由组
	group := r.Group("/group")
	group.Use(Logger3())
	{
		group.GET("/g-1", func(context *gin.Context) {
			println("作用于路由组....")
		})
	}
	r.Run(":8080")
}

func Logger1() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		context.Set("name", "tom")
		context.Next()
		latency := time.Since(t)
		log.Println("latency__1：", latency)
	}
}

func Logger2() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		context.Set("name", "tom")
		latency := time.Since(t)
		log.Println("latency__2：", latency)
		// 中间件的执行顺序就通过context.Next来控制
		context.Next()
	}
}

func Logger3() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		context.Set("name", "tom")
		context.Next()
		latency := time.Since(t)
		log.Println("latency__3：", latency)
	}
}
