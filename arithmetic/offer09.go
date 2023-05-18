package arithmetic

import "fmt"

/**
剑指Offer 09 用两个栈实现队列
*/

type Queue struct {
	stack1 []int
	stack2 []int
}

func Constructor() Queue {
	return Queue{[]int{}, []int{}}
}

func (thisQ *Queue) AppendTail(value int) {
	thisQ.stack1 = append(thisQ.stack1, value)
}

func (thisQ *Queue) DeleteHead() int {
	if len(thisQ.stack2) == 0 {
		for len(thisQ.stack1) > 0 {
			thisQ.stack2 = append(thisQ.stack2, thisQ.stack1[len(thisQ.stack1)-1])
			thisQ.stack1 = thisQ.stack1[:len(thisQ.stack1)-1]
		}
	}
	if len(thisQ.stack2) == 0 {
		return -1
	}
	res := thisQ.stack2[len(thisQ.stack2)-1]
	thisQ.stack2 = thisQ.stack2[:len(thisQ.stack2)-1]
	return res
}

func TestOffer09() {
	queue := Constructor()
	queue.AppendTail(5)
	queue.AppendTail(7)
	queue.AppendTail(3)
	fmt.Println(queue)
	head := queue.DeleteHead()
	fmt.Println(queue, head)
	queue.AppendTail(4)
	fmt.Println(queue)
	queue.DeleteHead()
	fmt.Println(queue)
}
