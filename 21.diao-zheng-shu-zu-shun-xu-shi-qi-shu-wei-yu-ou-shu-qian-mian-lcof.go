// [面试题21. 调整数组顺序使奇数位于偶数前面](https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/)

package main

import "fmt"

func exchange(nums []int) []int {
	var left, right int
	right = len(nums) - 1
	for left < right {
		for left < right && nums[left]%2 == 1 {
			left++
		}
		for left < right && nums[right]%2 == 0 {
			right--
		}

		if left != right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return nums
}

func exchangeMain() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
}
