// [面试题33. 二叉搜索树的后序遍历序列](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/)

package main

func verifyPostorder(postorder []int) bool {
	size := len(postorder)
	if size == 0 {
		return true
	}
	size--
	root := postorder[size]
	postorder = postorder[:size]

	var i int
	for i = 0; i < size; i++ {
		if postorder[i] > root {
			break
		}
	}
	for j := i + 1; j < size; j++ {
		if postorder[j] < root {
			return false
		}
	}
	result := true
	if i > 0 {
		left, right := postorder[:i-1], postorder[i:]
		result = verifyPostorder(left)
		result = result && verifyPostorder(right)
	} else {
		result = verifyPostorder(postorder)
	}
	return result
}
func verifyPostorderMain() {
	println(verifyPostorder([]int{4, 8, 6, 12, 16, 14, 10}), true)
	println(verifyPostorder([]int{4, 6, 12, 8, 16, 14, 10}), false)

	println(verifyPostorder([]int{1, 6, 3, 2, 5}), false)
	println(verifyPostorder([]int{1, 3, 2, 6, 5}), true)
}
