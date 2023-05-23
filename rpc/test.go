package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("父包中的 http 变量地址 %p\n", http.DefaultServeMux)
}
