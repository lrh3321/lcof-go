// [面试题46. 把数字翻译成字符串](https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/)

package main

import "strconv"

func translateNumStr(num string, memo map[string]int) int {
	if len(num) <= 1 {
		return 1
	}
	if v, exists := memo[num]; exists {
		return v
	}
	var result int

	first := num[0]
	switch first {
	case '1':
		result += translateNumStr(num[2:], memo)
	case '2':
		second := num[1]
		if second < '6' {
			result += translateNumStr(num[2:], memo)
		}
	}

	result += translateNumStr(num[1:], memo)

	memo[num] = result

	return result
}

func translateNum(num int) int {
	var memo = make(map[string]int)

	return translateNumStr(strconv.Itoa(num), memo)
}

func translateNumMain() {
	println(translateNum(25), 2)
	println(translateNum(624), 2)
	println(translateNum(6204), 2)
	println(translateNum(6214), 3)
	println(translateNum(2258), 3)
	println(translateNum(258), 2)
	println(translateNum(12258), 5)
}
