// [面试题61. 扑克牌中的顺子](https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/)

package main

import "sort"

func isStraight(nums []int) bool {
	sort.IntSlice(nums).Sort()
	jokerCount := 0
	for _, it := range nums {
		if it != 0 {
			break
		}
		jokerCount++
	}
	nums = nums[jokerCount:]
	if len(nums) == 0 {
		return true
	}
	pre := nums[0]
	for _, it := range nums[1:] {
		for pre+1 != it {
			if jokerCount <= 0 {
				return false
			}
			jokerCount--
			pre++
		}
		pre = it
	}
	return true
}

func isStraightMain() {
	println(isStraight([]int{1, 2, 3, 4, 5}), true)
	println(isStraight([]int{0, 0, 1, 2, 5}), true)
	println(isStraight([]int{11, 0, 9, 0, 0}), true)
	println(isStraight([]int{0, 0, 0, 1, 0}), true)
	println(isStraight([]int{0, 0, 0, 0, 0}), true)
	println(isStraight([]int{11, 0, 9, 1, 0}), false)
}
