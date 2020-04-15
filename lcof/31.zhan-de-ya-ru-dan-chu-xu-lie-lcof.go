// [面试题31. 栈的压入、弹出序列](https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/)

package main

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int

	for _, it := range pushed {
		stack = append(stack, it)
		if len(popped) == 0 {
			return true
		}

		top := popped[0]
		for len(stack) > 0 && stack[len(stack)-1] == top {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
			if len(popped) == 0 {
				return true
			}
			top = popped[0]
		}
	}

	return len(popped) == 0
}

func validateStackSequencesMain() {
	println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}), true)
	println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}), false)
}
