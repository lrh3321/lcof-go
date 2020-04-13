// [面试题03. 数组中重复的数字](https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/)
package main

func findRepeatNumber(nums []int) int {
	n := len(nums)
	bucket := make([]int, n, n)
	for _, it := range nums {
		if bucket[it] != 0 {
			return it
		}
		bucket[it]++
	}
	return 0
}

func findRepeatNumberMain() {
	println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}
