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

//func main() {
//	ch := make(chan int)
//	done := make(chan bool)
//	go producer(ch)
//	go customer(ch, done)
//	<-done
//	fmt.Println("All done!")
//	for i := 0; i < 3; i++ {
//		go downloadData("a.com/" + fmt.Sprint(i+'0'))
//	}
//	for i := 0; i < 3; i++ {
//		msg := <-chO
//		fmt.Println("finish", msg)
//	}
//	fmt.Println("Down done!")
//}

func add(num1 int, num2 int) int {
	return num1 + num2
}

type C struct {
	B
	money float32
}

type B struct {
	name string
	age  int
}

type TB interface {
	getNameByDay6() string
}

func (b B) getNameByDay6() string {
	return b.name
}
