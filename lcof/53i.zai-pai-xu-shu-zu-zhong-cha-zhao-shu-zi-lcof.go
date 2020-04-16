// [面试题53 - I. 在排序数组中查找数字 I](https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/)

package main

func search(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	left, right := 0, size-1
	found := false
	for left < right {
		mid := (left + right) >> 1
		it := nums[mid]
		if it == target {
			left = mid
			found = true
			break
		} else if it < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if !found {
		if left != right || nums[left] != target {
			return 0
		}
	}

	var result = 1
	for i := left + 1; i < size; i++ {
		if nums[i] != target {
			break
		}
		result++
	}
	for i := left - 1; i >= 0; i-- {
		if nums[i] != target {
			break
		}
		result++
	}

	return result
}

func searchMain() {
	println(search([]int{5, 7, 7, 8, 8, 10}, 8), 2)
	println(search([]int{5, 7, 7, 8, 8, 10}, 6), 0)
	println(search([]int{5, 7, 7, 9, 9, 10}, 6), 0)
}
