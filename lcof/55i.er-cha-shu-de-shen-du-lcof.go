// [面试题55 - I. 二叉树的深度](https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/)

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	d1, d2 := maxDepth(root.Left), maxDepth(root.Right)
	if d2 > d1 {
		d1 = d2
	}

	return d1 + 1
}

func maxDepthMain() {
	println(maxDepth(sliceToTree([]int{3, 9, 20, null, null, 15, 7})), 3)
}
