package main

import (
	"fmt"
	"net/http"
	"sync"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	server := &http.Server{
		Addr:    ":8081",
		Handler: http.HandlerFunc(helloHandler),
	}
	defer fmt.Println("结束了。。。。。111")
	fmt.Println("step：1")
	go func() {
		defer fmt.Println("结束了。。。。。")
		fmt.Println("step：3")
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("step：4")
			fmt.Println(err.Error())
		}
		fmt.Println("step：5")
	}()

	wg.Wait()
	fmt.Println("step：2")
}
