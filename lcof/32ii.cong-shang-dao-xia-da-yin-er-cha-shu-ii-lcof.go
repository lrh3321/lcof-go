// [面试题32 - II. 从上到下打印二叉树 II](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/)

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
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var queue = []*TreeNode{root}

	for len(queue) > 0 {
		q := queue
		queue = nil
		var row []int
		for _, it := range q {
			if it == nil {
				continue
			}
			queue = append(queue, it.Left, it.Right)
			row = append(row, it.Val)
		}
		if len(row) == 0 {
			break
		}
		result = append(result, row)
	}
	return result
}
func levelOrderMain() {
	fmt.Println(
		levelOrder(sliceToTree([]int{3, 9, 20, null, null, 15, 7})),
		[][]int{{3}, {9, 20}, {15, 7}},
	)
}
