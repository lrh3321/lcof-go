// [面试题56 - II. 数组中数字出现的次数 II](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/)

package main

func singleNumber(nums []int) int {
	var result int
	for i := 0; i < 32; i++ {
		mask := 1 << i
		count := 0
		for _, it := range nums {
			if mask&it != 0 {
				count++
			}
		}
		if count%3 != 0 {
			result |= mask
		}
	}

	return result
}

func singleNumberMain() {
	println(singleNumber([]int{3, 4, 3, 3}), 4)
	println(singleNumber([]int{9, 1, 7, 9, 7, 9, 7}), 1)
}
