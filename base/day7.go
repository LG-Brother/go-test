package main

import (
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<H1>详情</H1></br>hello world"))
}

//
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(1)
//	mux := http.NewServeMux()
//	mux.HandleFunc("/hello", helloHandler)
//	mux.HandleFunc("/info", infoHandler)
//	server := &http.Server{
//		Addr:    ":8081",
//		Handler: mux,
//	}
//	defer fmt.Println("结束了。。。。。111")
//	fmt.Println("step：1")
//	go func() {
//		defer fmt.Println("结束了。。。。。")
//		fmt.Println("step：3")
//		defer wg.Done()
//		if err := server.ListenAndServe(); err != nil {
//			fmt.Println("step：4")
//			fmt.Println(err.Error())
//		}
//		fmt.Println("step：5")
//	}()
//
//	wg.Wait()
//	fmt.Println("step：2")
//}

//func main() {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/hello", helloHandler)
//	mux.HandleFunc("/info", infoHandler)
//	server := &http.Server{
//		Addr:    ":8081",
//		Handler: mux,
//	}
//	server.ListenAndServe()
//}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/info", infoHandler)
	server := &http.Server{
		Addr: ":8081",
	}
	server.ListenAndServe()
}
