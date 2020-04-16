// [面试题53 - II. 0～n-1中缺失的数字](https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/)

package main

func missingNumber(nums []int) int {
	size := len(nums)
	if nums[size-1] != size {
		return size
	}

	left, right := 0, size-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if nums[left] == left {
		return left + 1
	}

	return left
}

func missingNumberMain() {
	println(missingNumber([]int{0, 2, 3}), 1)

	println(missingNumber([]int{0, 1, 3}), 2)
	println(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}), 8)
	println(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}), 9)
}
