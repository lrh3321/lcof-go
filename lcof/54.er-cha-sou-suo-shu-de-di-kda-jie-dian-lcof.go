// [面试题54. 二叉搜索树的第k大节点](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/)

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthLargestDFS(root *TreeNode, k, result *int) {
	if root != nil {
		kthLargestDFS(root.Right, k, result)
		(*k)--
		if *k == 0 {
			*result = root.Val
			return
		}
		kthLargestDFS(root.Left, k, result)
	}
}

func kthLargest(root *TreeNode, k int) int {

	var result int
	kthLargestDFS(root, &k, &result)
	return result
}

func kthLargestMain() {
	println(
		kthLargest(sliceToTree([]int{3, 1, 4, null, 2}), 1),
		4,
	)

	println(
		kthLargest(sliceToTree([]int{5, 3, 6, 2, 4, null, null, 1}), 3),
		4,
	)
}
