package arithmetic

import "fmt"

/**
剑指Offer 07 重建二叉树
*/

// 递归方法
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 通过前序遍历找到根节点
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree1(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree1(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

// 迭代法
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	// 定义临时栈
	var stack []*TreeNode
	stack = append(stack, root)
	// 定义中序遍历指针
	inorderIndex := 0
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			// 弹栈并移动中序遍历的指针
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}

func TestBuildTree() {
	preorder := []int{3, 9, 8, 5, 4, 10, 20, 15, 7}
	inorder := []int{4, 5, 8, 10, 9, 3, 15, 20, 7}
	tree := buildTree2(preorder, inorder)
	fmt.Println(PreorderPrintTreeByRecursion(tree))
}
