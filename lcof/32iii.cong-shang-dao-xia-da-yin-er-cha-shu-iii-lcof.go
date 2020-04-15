// [面试题32 - III. 从上到下打印二叉树 III](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/)

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
	reverse := false
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
		size := len(row)
		if size == 0 {
			break
		}
		if reverse {
			half := size >> 1
			for i := 0; i < half; i++ {
				row[i], row[size-1-i] = row[size-1-i], row[i]
			}
		}
		reverse = !reverse
		result = append(result, row)
	}
	return result
}
func levelOrderMain() {
	fmt.Println(
		levelOrder(sliceToTree([]int{3, 9, 20, null, null, 15, 7})),
		[][]int{{3}, {20, 9}, {15, 7}},
	)
}
