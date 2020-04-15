// [面试题11. 旋转数组的最小数字](https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/)

package main

import "fmt"

func minArray(numbers []int) int {
	size := len(numbers)
	if size == 1 {
		return numbers[0]
	}
	minIndex, maxIndex := size-1, 0

	min, max := numbers[minIndex], numbers[maxIndex]
	if min > max {
		return max
	}
	i := (minIndex + maxIndex) >> 1
	if min == max {
		m1, m2 := minArray(numbers[:i+1]), minArray(numbers[i+1:])
		if m2 < m1 {
			return m2
		}
		return m1
	}

	for maxIndex+1 != minIndex {
		it := numbers[i]
		if it >= max {
			maxIndex = i
			max = it
		} else if it <= min {
			minIndex = i
			min = it
		}
		i = (minIndex + maxIndex) >> 1
	}

	return min
}

func minArrayMain() {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}), 1)
	fmt.Println(minArray([]int{2, 2, 2, 0, 1}), 0)
	fmt.Println(minArray([]int{1, 3, 5}), 1)
	fmt.Println(minArray([]int{9}), 9)
	fmt.Println(minArray([]int{9, 9, 9, 9}), 9)
}
