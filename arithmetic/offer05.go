package arithmetic

import (
	"fmt"
)

/**
剑指Offer 05 替换空格
*/

// 遍历替换
func replace1(s string) string {
	res := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res = append(res, []byte("%20")...)
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}

// 原地替换
func replace2(s string) string {

	return ""
}

func TestReplace() {
	str := "He l lo Wo l d"
	fmt.Println("替换前:", str)
	//str = strings.Replace(str, " ", "%20", -1)
	str = replace1(str)
	fmt.Println("替换前:", str)
}
