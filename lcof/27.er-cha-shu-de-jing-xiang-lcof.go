// [面试题27. 二叉树的镜像](https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/)

package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := mirrorTree(root.Right)
	right := mirrorTree(root.Left)
	root.Left, root.Right = left, right
	return root
}

func mirrorTreeMain() {

	fmt.Println(
		treeToArray(
			mirrorTree(
				sliceToTree([]int{4, 2, 7, 1, 3, 6, 9}),
			)),
		[]int{4, 7, 2, 9, 6, 3, 1},
	)

	fmt.Println(
		treeToArray(
			mirrorTree(
				sliceToTree([]int{4, 2, 7, 1, 3, 6}),
			)),
		[]interface{}{4, 7, 2, nil, 6, 3, 1},
	)
}
