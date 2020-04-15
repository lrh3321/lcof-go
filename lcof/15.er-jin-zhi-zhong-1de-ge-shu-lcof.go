// [面试题15. 二进制中1的个数](https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/)

package main

func hammingWeight(num uint32) int {
	var result int

	for num > 0 {
		result++
		num &= num - 1
	}
	return result
}

func hammingWeightMain() {
	println(hammingWeight(9), 2)
	println(hammingWeight(35), 3)
}
