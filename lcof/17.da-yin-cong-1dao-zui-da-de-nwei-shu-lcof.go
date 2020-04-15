// [面试题17. 打印从1到最大的n位数](https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/)

package main

func printNumbers(n int) []int {
	cap := 1
	for i := 0; i < n; i++ {
		cap *= 10
	}
	cap--
	result := make([]int, cap, cap)

	for i := 0; i < cap; i++ {
		result[i] = i + 1
	}

	return result
}
