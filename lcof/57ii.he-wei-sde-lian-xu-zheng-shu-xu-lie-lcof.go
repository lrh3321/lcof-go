// [面试题57 - II. 和为s的连续正数序列](https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/)

package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	var result [][]int
	t2 := target << 1

	for i := 2; i < target; i++ {
		var n int

		if i&1 == 0 {
			if t2%i != 0 {
				continue
			}

			k := target / i
			n = k - (i >> 1) + 1
		} else {
			if target%i != 0 {
				continue
			}
			k := target / i
			n = k - (i >> 1)
		}
		if n < 1 {
			break
		}
		tmp := ((n << 1) + i - 1) * i
		if tmp != t2 {
			continue
		}
		t := make([]int, i, i)
		for j := 0; j < i; j++ {
			t[j] = j + n
		}
		result = append(result, t)
	}

	left, right := 0, len(result)-1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}

	return result
}

func findContinuousSequenceMain() {
	fmt.Println(findContinuousSequence(9), [][]int{{2, 3, 4}, {4, 5}})
	fmt.Println(findContinuousSequence(10), [][]int{{1, 2, 3, 4}})
	fmt.Println(findContinuousSequence(15), [][]int{{1, 2, 3, 4, 5}, {4, 5, 6}, {7, 8}})
}
