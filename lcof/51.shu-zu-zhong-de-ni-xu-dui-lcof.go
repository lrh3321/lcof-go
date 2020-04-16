// [面试题51. 数组中的逆序对](https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/)

package main

var reversePairsTemp = make([]int, 50000)

func reversePairs(nums []int) int {
	size := len(nums)
	if size < 2 {
		return 0
	}
	if size == 2 {
		if nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
			return 1
		}
		return 0
	}
	var result int
	half := size >> 1
	left, right := nums[:half], nums[half:]

	result += reversePairs(left)
	result += reversePairs(right)
	pos := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			reversePairsTemp[pos] = left[0]
			left = left[1:]
		} else {
			reversePairsTemp[pos] = right[0]
			right = right[1:]
			result += len(left)
		}
		pos++
	}

	copy(reversePairsTemp[pos:], left)
	copy(reversePairsTemp[pos:], right)
	copy(nums, reversePairsTemp[:size])

	return result
}

func reversePairsMain() {
	var nums []int
	nums = []int{1, 3, 2, 3, 1}
	println(reversePairs(nums), 5)
	nums = []int{7, 5, 6, 4}
	println(reversePairs(nums), 5)
}
