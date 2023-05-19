package main

import (
	"fmt"
	"runtime"
)

var title = "test go"

func getAge() int {
	return 25
}

func getName() string {
	return "LinGang"
}

func routine1() {
	println("this is routine 1")
	printRoutineInfo()
}

func routine2() {
	println("this is routine 2")
	printRoutineInfo()
}

func printRoutineInfo() {
	buf := make([]byte, 2048)
	stackInfo := runtime.Stack(buf, true)
	fmt.Printf("====Begin stack tracke =====\n%v\n====End stack tracke====\n", stackInfo)
}

//func main() {
//	printRoutineInfo()
//	name := getAge()
//	age := getName()
//	go routine1()
//	go routine2()
//	fmt.Println(title, name, age)
//	time.Sleep(10000)
//}
