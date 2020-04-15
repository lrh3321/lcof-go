// [面试题42. 连续子数组的最大和](https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/)

package main

func maxSubArray(nums []int) int {
	size := len(nums)
	dp := make([]int, size, size)
	copy(dp, nums)
	var max = nums[0]
	for i := 1; i < size; i++ {
		t := dp[i-1]
		if t > 0 {
			dp[i] += t
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

func maxSubArrayMain() {
	println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6)

	println(maxSubArray([]int{1, 2}), 3)
}
