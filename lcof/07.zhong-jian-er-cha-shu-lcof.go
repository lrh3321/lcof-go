// [面试题07. 重建二叉树](https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/)

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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	first := preorder[0]
	root := &TreeNode{Val: first}
	var i int
	for index, it := range inorder {
		if it == first {
			i = index
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

func buildTreeMain() {
	fmt.Println(treeToArray(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})))
}
