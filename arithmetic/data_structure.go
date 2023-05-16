package arithmetic

import (
	"math/rand"
	"time"
)

/**
定义数据结构
*/

// ListNode 链表结构
type ListNode struct {
	Value int
	Next  *ListNode
}

// CreateLinkedList 创建指定个数的随机链表
func CreateLinkedList(count int) *ListNode {
	head := &ListNode{}
	tail := head
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		node := &ListNode{Value: rand.Intn(100)}
		tail.Next = node
		tail = node
	}
	return head
}

func PrintLinkedList(head *ListNode) {
	if head == nil {
		return
	}
	p := head.Next
	for p != nil {
		print(p.Value, " ")
		p = p.Next
	}
	println()
}
