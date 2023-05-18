package arithmetic

import (
	"fmt"
	"time"
)

/**
剑指Offer 10-1 斐波那契数列
*/

// 方法一：递归法
func fib1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

// 方法二：记忆化递归法
func fib2(n int) int {
	// 创建记忆数组
	arr := make([]int, n+2)
	arr[0] = 0
	arr[1] = 1
	// 递归运算
	var recursion func(int) int
	recursion = func(i int) int {
		if i > 1 && arr[i] == 0 {
			arr[i] = (recursion(i-1) + recursion(i-2)) % 1000000007
		}
		return arr[i]
	}
	recursion(n)
	return arr[n]
}

// 方法三：动态规划 迭代法
func fib3(n int) int {
	if n < 2 {
		return n
	}
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = (arr[i-1] + arr[i-2]) % 1000000007
	}
	return arr[n]
}

func TestFib() {
	n := 45
	start := time.Now()
	fmt.Println("输入数据:", n)
	fmt.Println("输出结果:", fib3(n))
	elapsed := time.Since(start)
	fmt.Println("计算耗时:", elapsed)
}
