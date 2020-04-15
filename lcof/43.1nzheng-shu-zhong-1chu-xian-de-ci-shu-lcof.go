// [面试题43. 1～n整数中1出现的次数](https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/)

package main

func countDigitOne(n int) int {
	var count int
	i := 1
	for n/i != 0 {
		high := n / (i * 10)
		cur := n / i % 10
		low := n % i

		if cur == 0 {
			count += high * i
		} else if cur == 1 {
			count += high*i + low + 1
		} else {
			count += high*i + i
		}

		i *= 10
	}

	return count
}

func countDigitOneMain() {
	println(countDigitOne(12), 5)
	println(countDigitOne(13), 6)
}
