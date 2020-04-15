// [面试题67. 把字符串转换成整数](https://leetcode-cn.com/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/)

package main

import "strings"

func strToInt(str string) int {
	str = strings.TrimSpace(str)
	p := []byte(str)
	var result int64
	// hasSign := false
	isMinus := false
For:
	for i, c := range p {
		switch {
		case c == '-':
			if i != 0 {
				break For
			}
			isMinus = true
		case c == '+':
			if i != 0 {
				break For
			}
		case c < '0' || c > '9':
			break For
		default:
			result = result*10 + int64(c-'0')
			if result > 0x7fffffff {
				break For
			}
		}
	}
	if isMinus {
		if result >= 0x80000000 {
			return -2147483648
		}
		result = -result
	} else if result > 0x7fffffff {
		return 0x7fffffff
	}
	return int(result)
}

func strToIntMain() {
	println(strToInt("42"), 42)
	println(strToInt("   -42"), -42)
	println(strToInt("4193 with words"), 4193)
	println(strToInt("-91283472332"), -2147483648)
	println(strToInt("91283472332"), 2147483647)
	println(strToInt("123-"), 123)
}
