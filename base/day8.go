package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", day8Handler)
	http.ListenAndServe(":8080", nil)
}

func day8Handler(writer http.ResponseWriter, request *http.Request) {
	method := request.Method
	if http.MethodPost == method {
		log.Println("This is [Post] request")
	} else if http.MethodGet == method {
		log.Println("This is [Get] request")
	} else {
		log.Println("This is [other] request")
	}
}
