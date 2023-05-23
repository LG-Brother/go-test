package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

type Cal int

func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	log.Println("Square was Called")
	return nil
}

func main() {
	rpc.Register(new(Cal))
	rpc.HandleHTTP()
	log.Printf("Serving RPC server on port %d", 1234)
	fmt.Printf("子包中的 http 变量地址 %p\n", http.DefaultServeMux)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Error serving", err)
	}
}
