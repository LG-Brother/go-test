package arithmetic

import "fmt"

/*
*
剑指Offer 06 从头到尾打印链表
*/

// 顺序遍历 + 反转
func reversPrint1(head *ListNode) []int {
	var ans []int
	if head == nil {
		return ans
	}
	p := head.Next
	for ; p != nil; p = p.Next {
		ans = append(ans, p.Value)
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

// 递归法
func reversPrint2(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	ans := reversPrint2(head.Next)
	ans = append(ans, head.Value)
	return ans
}

func TestReversPrint() {
	head := CreateLinkedList(5)
	PrintLinkedList(head)
	ans := reversPrint2(head)
	fmt.Println(ans)
}
