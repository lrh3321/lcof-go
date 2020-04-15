// [面试题34. 二叉树中和为某一值的路径](https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/)

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
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	sum -= root.Val

	if root.Left == nil && root.Right == nil {
		// 子叶节点
		if sum == 0 {
			return [][]int{{root.Val}}
		}
	}
	var result [][]int
	for _, subTree := range []*TreeNode{root.Left, root.Right} {
		for _, it := range pathSum(subTree, sum) {
			row := make([]int, len(it)+1)
			row[0] = root.Val
			copy(row[1:], it)
			result = append(result, row)
		}
	}

	return result
}

func pathSumMain() {
	fmt.Println(
		pathSum(sliceToTree([]int{
			5,
			4, 8,
			11, nullTreeValue, 13, 4,
			7, 2, nullTreeValue, nullTreeValue, nullTreeValue, nullTreeValue, 5, 1,
		}), 22),
		[][]int{
			{5, 4, 11, 2},
			{5, 8, 4, 5},
		},
	)
	fmt.Println(
		pathSum(sliceToTree([]int{
			-2,
			nullTreeValue, -3,
		}), -5),
		[][]int{
			{-2, -3},
		},
	)
	fmt.Println(
		pathSum(sliceToTree([]int{
			1,
			-2, -3,
			1, 3, -2, nullTreeValue,
			-1,
		}), -1),
		[][]int{
			{1, -2, 1, -1},
		},
	)
}
