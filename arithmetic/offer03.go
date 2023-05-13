package arithmetic

import "fmt"

/**
剑指Offer 03 找出数组中重复的数字
*/

// 利用Hash表
func findRepeatNumber1(arr []int) int {
	vis := map[int]bool{}
	for _, num := range arr {
		if vis[num] {
			return num
		}
		vis[num] = true
	}
	return -1
}

// 原地交换
func findRepeatNumber2(arr []int) int {
	for i := 0; i < len(arr); {
		fmt.Println("i=", i, "打印数组：", arr)
		if arr[i] == i {
			i++
			continue
		}
		if arr[i] == arr[arr[i]] {
			return arr[i]
		}
		arr[arr[i]], arr[i] = arr[i], arr[arr[i]]
	}
	return -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func TestFindRepeatNumber() {
	arr := []int{3, 6, 0, 1, 2, 5, 3}
	fmt.Println("重复的数字为：", findRepeatNumber2(arr))
}
