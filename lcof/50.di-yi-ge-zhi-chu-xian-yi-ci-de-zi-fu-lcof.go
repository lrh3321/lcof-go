// [面试题50. 第一个只出现一次的字符](https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/)

package main

func firstUniqChar(s string) byte {
	if s == "" {
		return ' '
	}
	p := []byte(s)
	bucket := make([]int, 256, 256)
	var charSeq []byte
	for _, c := range p {
		if bucket[c] == 0 {
			charSeq = append(charSeq, c)
		}
		bucket[c]++
	}

	for _, c := range charSeq {
		if bucket[c] == 1 {
			return c
		}
	}
	return ' '
}

func firstUniqCharMain() {
	println(firstUniqChar("abaccdeff"), 'b')
	println(firstUniqChar(""), ' ')
}
