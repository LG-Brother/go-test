package main

type Test interface {
	getPName() string
	getMoney() float32
}

type Person interface {
	getPName() string
	getPAge() int
}

type Teacher struct {
	name string
	age  int
}

func (tea *Teacher) getPName() string {
	return tea.name
}

func (tea *Teacher) getPAge() int {
	return tea.age
}

type Worker struct {
	name  string
	age   int
	money float32
}

func (wor *Worker) getMoney() float32 {
	return 100.0
}

func (wor *Worker) getPName() string {
	return "老六"
}

func (wor *Worker) getPAge() int {
	return wor.age
}

//func main() {
//	var p Test
//	worker := Worker{name: "tom", age: 18}
//	p = &worker
//	fmt.Println(p.getPName())
//}
