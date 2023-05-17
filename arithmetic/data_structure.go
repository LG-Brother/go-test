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
	Val  int
	Next *ListNode
}

// TreeNode 二叉树结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// CreateLinkedList 创建指定个数的随机链表
func CreateLinkedList(count int) *ListNode {
	head := &ListNode{}
	tail := head
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		node := &ListNode{Val: rand.Intn(100)}
		tail.Next = node
		tail = node
	}
	return head
}

// PrintLinkedList 打印链表
func PrintLinkedList(head *ListNode) {
	if head == nil {
		return
	}
	p := head.Next
	for p != nil {
		print(p.Val, " ")
		p = p.Next
	}
	println()
}

// PreorderPrintTree 使用非递归方法先序遍历打印二叉树
func PreorderPrintTree(root *TreeNode) []int {

	return nil
}

// PreorderPrintTreeByRecursion 使用递归方法先序遍历打印二叉树
func PreorderPrintTreeByRecursion(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	left := PreorderPrintTreeByRecursion(root.Left)
	right := PreorderPrintTreeByRecursion(root.Right)
	ans = append(ans, root.Val)
	ans = append(ans, left...)
	ans = append(ans, right...)
	return ans
}

// LevelOrder 二叉树进行层序遍历
func LevelOrder(root *TreeNode) []int {

	return nil
}
