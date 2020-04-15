// [面试题66. 构建乘积数组](https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/)

package main

import "fmt"

func constructArr(a []int) []int {
	size := len(a)
	if size == 0 {
		return nil
	}
	left, right, result := make([]int, size, size), make([]int, size, size), make([]int, size, size)
	left[0] = 1
	right[size-1] = 1
	for i := 1; i < size; i++ {
		left[i] = left[i-1] * a[i-1]
	}
	for i := size - 2; i >= 0; i-- {
		right[i] = right[i+1] * a[i+1]
	}

	for i := 0; i < size; i++ {
		result[i] = left[i] * right[i]
	}
	return result
}

func constructArrMain() {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}), []int{120, 60, 40, 30, 24})
}
