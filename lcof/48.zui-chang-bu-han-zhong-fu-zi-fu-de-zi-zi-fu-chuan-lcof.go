// [面试题48. 最长不含重复字符的子字符串](https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/)

package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	p := []byte(s)
	size := len(p)
	bucket := make([]int, 256, 256)
	bucket[p[0]] = 1
	left, right := 0, 1
	var result = 1

	for left <= right && right < size {

		for right < size && bucket[p[right]] == 0 {
			bucket[p[right]]++
			right++
		}

		t := right - left
		if t > result {
			result = t
		}

		if right == size {
			return result
		}
		c := p[right]
		for left <= right && p[left] != c {
			bucket[p[left]]--
			left++
		}
		bucket[p[left]]--
		left++
	}

	return result
}

func lengthOfLongestSubstringMain() {
	println(lengthOfLongestSubstring("abcabcbb"), 3)
	println(lengthOfLongestSubstring("bbbbb"), 1)
	println(lengthOfLongestSubstring("pwwkew"), 3)
}
