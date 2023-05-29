package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name1", "test1", "test2")
	age := flag.String("age1", "18", "test")
	flag.Parse()
	fmt.Println("name ->", name)
	fmt.Println("age ->", *age)
}
