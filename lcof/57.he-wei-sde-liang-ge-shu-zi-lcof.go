// [面试题57. 和为s的两个数字](https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/)

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	size := len(nums)
	if size < 2 {
		return nil
	}
	left, right := 0, size-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{nums[left], nums[right]}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

func twoSumMain() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9), []int{2, 7})
	fmt.Println(twoSum([]int{10, 26, 30, 31, 47, 60}, 40), []int{10, 30})
}
