// [面试题56 - I. 数组中数字出现的次数](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/)

package main

import "fmt"

func singleNumbers(nums []int) []int {
	var xor, a, b int
	for _, it := range nums {
		xor ^= it
	}
	mask := 1
	for xor&mask == 0 {
		mask <<= 1
	}
	for _, it := range nums {
		if mask&it == 0 {
			a ^= it
		} else {
			b ^= it
		}

	}

	return []int{a, b}
}

func singleNumbersMain() {
	fmt.Println(singleNumbers([]int{4, 1, 4, 6}), []int{1, 6})
	fmt.Println(singleNumbers([]int{1, 2, 10, 4, 1, 4, 3, 3}), []int{2, 10})
}
