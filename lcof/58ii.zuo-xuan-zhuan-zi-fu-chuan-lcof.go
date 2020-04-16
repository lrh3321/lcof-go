// [面试题58 - II. 左旋转字符串](https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/)

package main

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

func reverseLeftWordsMain() {
	println(
		reverseLeftWords("abcdefg", 2),
		"cdefgab",
	)
	println(
		reverseLeftWords("lrloseumgh", 6),
		"umghlrlose",
	)
}
