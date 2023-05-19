package main

import "fmt"

const (
	MALE   = 1
	FEMALE = 0
)

type student struct {
	name string
	age  int
}

//func main() {
//	sex := 1
//	switch sex {
//	case MALE:
//		fmt.Println("It's a Man")
//	case FEMALE:
//		fmt.Println("It's a Woman")
//	}
//	s1 := student{name: "tom"}
//	s1.age = 24
//	s2 := new(student)
//	s2.name = "lucy"
//	s2.age = 25
//	fmt.Println("【改名前-------】")
//	fmt.Println(s1)
//	fmt.Println(s2)
//	changeName1(s1, "大傻逼")
//	changeName2(s2, "臭流氓")
//	fmt.Println("【改名后-------】")
//	fmt.Println(s1)
//	fmt.Println(s2)
//}

func changeName1(stu student, name string) {
	stu.name = name
	fmt.Println("changeName1-改名咯:", stu)
}
func changeName2(stu *student, name string) {
	stu.name = name
	fmt.Println("changeName2-改名咯:", stu)
}
