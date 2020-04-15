// [面试题28. 对称的二叉树](https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/)

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	n := 2
	half := n >> 1
	row := []*TreeNode{root.Left, root.Right}
	loop := true
	for loop {
		loop = false
		nextRow := make([]*TreeNode, n*2, n*2)
		for i := 0; i < half; i++ {
			j := n - i - 1
			t1, t2 := row[i], row[j]
			if t1 == nil && t2 == nil {
			} else if t1 != nil && t2 != nil && t1.Val == t2.Val {
				nextRow[2*i], nextRow[2*i+1] = t1.Left, t1.Right
				nextRow[2*j], nextRow[2*j+1] = t2.Left, t2.Right

				if !loop && (t1.Left != nil || t1.Right != nil || t2.Left != nil || t2.Right != nil) {
					loop = true
				}
			} else {
				return false
			}
		}
		n <<= 1
		half <<= 1
		row = nextRow
	}
	return true
}

func isSymmetricMain() {
	println(
		isSymmetric(
			sliceToTree([]int{1, 2, 2, 3, 4, 4, 3}),
		),
		true,
	)
	println(
		isSymmetric(
			sliceToTree([]int{1, 2, 2, nullTreeValue, 3, nullTreeValue, 3}),
		),
		false,
	)
}
