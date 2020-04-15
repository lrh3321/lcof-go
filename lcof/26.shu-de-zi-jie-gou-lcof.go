// [面试题26. 树的子结构](https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/)

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructureImpl(A *TreeNode, B *TreeNode, nested bool) bool {
	if A == nil {
		return false
	}

	if B == nil {
		return false
	}

	if A.Val == B.Val {
		result := true
		if result && A.Left == nil && B.Left != nil {
			result = false
		}
		if result && A.Right == nil && B.Right != nil {
			result = false
		}

		if result && A.Left != nil && B.Left != nil {
			result = isSubStructureImpl(A.Left, B.Left, true)
		}
		if result && A.Right != nil && B.Right != nil {
			result = isSubStructureImpl(A.Right, B.Right, true)
		}

		if result {
			return true
		}
	}
	if nested {
		return false
	}
	return isSubStructureImpl(A.Left, B, false) || isSubStructureImpl(A.Right, B, false)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return isSubStructureImpl(A, B, false)
}

func isSubStructureMain() {
	// [1,2,3]
	// [3,1]
	println(isSubStructure(
		sliceToTree([]int{1, 2, 3}),
		sliceToTree([]int{3, 1}),
	), false)
	// [3,4,5,1,2]
	// [4,1]
	println(isSubStructure(
		sliceToTree([]int{3, 4, 5, 1, 2}),
		sliceToTree([]int{4, 1}),
	), true)
	// [1,0,1,-4,-3]
	// [1,-4]
	println(isSubStructure(
		sliceToTree([]int{1, 0, 1, -4, -3}),
		sliceToTree([]int{1, -4}),
	), false)
}
