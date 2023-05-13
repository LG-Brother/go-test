package arithmetic

import (
	"fmt"
	"sort"
)

/**
剑指Offer 04 二维数组中的查找
*/

// 二分查找一维数组
func binarySearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 二分法查找
func findNumIn2DArray1(matrix [][]int, target int) bool {
	for _, row := range matrix {
		j := sort.SearchInts(row, target)
		if j < len(matrix[0]) && row[j] == target {
			return true
		}
	}
	return false
}

// 从右上角或左下角搜索
func findNumIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}

	}
	return false
}

func createMatrix() [][]int {
	matrix := make([][]int, 3)
	fmt.Println("初始化二维数组：\n", matrix)
	for i := range matrix {
		matrix[i] = make([]int, 4)
	}
	fmt.Println("初始化内部的一维数组：\n", matrix)
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = i*10 + j
		}
	}
	fmt.Println("二维数组赋值：\n", matrix)
	return matrix
}

func TestFindNumIn2DArray() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println("二维数组为：\n", matrix)
	target := 26
	fmt.Printf("二维数组中寻找%v\n", target)
	if findNumIn2DArray2(matrix, target) {
		fmt.Println("------> 已找到")
	} else {
		fmt.Println("------> 未找到")
	}
}
