package main

import (
	"fmt"
	"time"
)

// 生产者
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

// 消费者
func customer(ch <-chan int, done chan<- bool) {
	for num := range ch {
		fmt.Println("Received:", num)
	}
	done <- true
}

var chO = make(chan string, 10)

func downloadData(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	chO <- url
}

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go producer(ch)
	go customer(ch, done)
	<-done
	fmt.Println("All done!")
	for i := 0; i < 3; i++ {
		go downloadData("a.com/" + string(i+'0'))
	}
	for i := 0; i < 3; i++ {
		msg := <-chO
		fmt.Println("finish", msg)
	}
	fmt.Println("Down done!")
}
