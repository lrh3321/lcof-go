// [面试题55 - II. 平衡二叉树](https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/)

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func treeDepth(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	d1, delta1 := treeDepth(root.Left)
	if delta1 > 1 {
		return 0, delta1
	}
	d2, delta2 := treeDepth(root.Right)
	if delta2 > 1 {
		return 0, delta2
	}

	if d2 > d1 {
		delta1 = d2 - d1
		d1 = d2
	} else {
		delta1 = d1 - d2
	}

	return d1 + 1, delta1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, delta := treeDepth(root)

	return delta < 2
}

func isBalancedMain() {
	println(isBalanced(sliceToTree([]int{3, 9, 20, null, null, 15, 7})), true)
	println(isBalanced(sliceToTree([]int{1, 2, 2, 3, 3, null, null, 4, 4})), false)
}
