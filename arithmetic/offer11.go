package arithmetic

import "fmt"

/**
剑指Offer 11 旋转数组中的最小数字
描述：
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
给你一个可能存在 重复 元素值的数组 numbers ，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。请返回旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为 1。
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
*/

func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1
	for l < r {
		mid := (l + r) / 2
		if numbers[mid] > numbers[r] {
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			r = mid
		} else {
			r--
		}
	}
	return numbers[l]
}

func TestMinArray() {
	arr := []int{3, 4, 5, 1, 2}
	fmt.Println("数组中最小值为:", minArray(arr))
}
