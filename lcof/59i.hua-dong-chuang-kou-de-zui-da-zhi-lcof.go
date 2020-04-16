// [面试题59 - I. 滑动窗口的最大值](https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/)

package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if k == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}

	var result []int
	var queue = []int{nums[0]}
	for i := 1; i < k; i++ {
		it := nums[i]

		for len(queue) > 0 && it > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, it)
	}
	result = append(result, queue[0])

	for i, it := range nums[k:] {
		poped := nums[i]
		top := queue[0]
		if poped == top {
			queue = queue[1:]
			if len(queue) > 0 {
				top = queue[0]
			}
		}

		for len(queue) > 0 && it > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, it)

		result = append(result, queue[0])
	}

	return result
}

func maxSlidingWindowMain() {
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3), []int{3, 3, 2, 5})
	fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2), []int{7, 4})

	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3), []int{3, 3, 5, 5, 6, 7})
	fmt.Println(maxSlidingWindow([]int{}, 0), []int{})
	fmt.Println(maxSlidingWindow([]int{1, -1}, 1), []int{1, -1})
}
