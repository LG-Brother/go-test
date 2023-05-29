package main

import (
	"fmt"
	"reflect"
)

func main() {
	idTag := reflect.TypeOf(BlogTag{}).Field(0).Tag.Get("json")
	fmt.Println("Tag for Id filed:", idTag)
	nameTag := reflect.TypeOf(BlogTag{}).Field(1).Tag.Get("test")
	fmt.Println("Tag for Name filed:", nameTag)
	value, ok := reflect.TypeOf(BlogTag{}).Field(1).Tag.Lookup("test1")
	fmt.Println("ok:", ok, ", value:", value)
}

type BlogTag struct {
	Id   rune   `json:"id"`
	Name string `json:"name" test:"testTag"`
}
