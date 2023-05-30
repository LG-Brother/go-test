package arithmetic

import "fmt"

/**
剑指Offer 12 矩阵中的路径
*/

func exist(board [][]byte, word string) bool {
	// 避免取值出错
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		// 先临时修改该位置的值，避免搜索下一个值时造成影响
		board[i][j] = ' '
		// 按照下、右、左、上顺序搜索下一值
		ans := dfs(i+1, j, k+1) || dfs(i, j+1, k+1) || dfs(i, j-1, k+1) || dfs(i-1, j, k+1)
		// 值复原
		board[i][j] = word[k]
		return ans
	}
	for i, arr := range board {
		for j, _ := range arr {
			fmt.Println("i:", i, ", j:", j)
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func TestExist() {
	board := [][]byte{
		{'C', 'A', 'A'},
		{'A', 'A', 'A'},
		{'B', 'C', 'D'},
	}
	word := "AAB"
	fmt.Println("矩阵是否存在路径：", exist(board, word))
}
