package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("*/3 * * * * *", func() {
		fmt.Println("每秒执行一次:", time.Now().Format("2006-01-02 15:04:05"))
	})
	if err != nil {
		fmt.Println("添加任务失败:", err)
		return
	}
	c.Start()
	println("当前时间为：", time.Now().Format("2006-01-02 15:04:05"))
	select {}
}
