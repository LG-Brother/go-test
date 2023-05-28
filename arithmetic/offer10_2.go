package arithmetic

import "fmt"

/**
剑指Offer 10-2 青蛙跳台阶问题
描述：
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
*/

// 递归法
func numWays1(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return numWays1(n-2) + numWays1(n-1)
}

// 动态规划 迭代法
func numWay2(n int) int {
	if n == 1 {
		return 1
	}
	arr := make([]int, n+1)
	arr[0] = 1
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = (arr[i-1] + arr[i-2]) % 1000000007
	}
	return arr[n]
}

func TestNumWays() {
	fmt.Print("请输入台阶数：")
	var num int
	fmt.Scanln(&num)
	fmt.Printf(" %v 级台阶总共有 %v 中跳法", num, numWay2(num))
}
