// [面试题39. 数组中出现次数超过一半的数字](https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/)

package main

import "sort"

func majorityElement(nums []int) int {
	sort.IntSlice(nums).Sort()
	half := len(nums) >> 1

	return nums[half]
}
