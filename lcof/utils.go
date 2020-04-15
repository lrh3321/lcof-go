package main

import "strconv"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return "null"
	}
	return strconv.Itoa(l.Val)
}

func sliceToListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var head = &ListNode{Val: arr[0]}
	cur := head
	for _, it := range arr[1:] {
		next := &ListNode{Val: it}
		cur.Next = next
		cur = next
	}

	return head
}

func listNodeToSlice(l *ListNode) []*ListNode {
	var result []*ListNode
	cur := l
	for cur != nil {
		result = append(result, cur)
		cur = cur.Next
	}
	return result
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	if t == nil {
		return "<nil>"
	}
	return strconv.Itoa(t.Val)
}

func treeToArray(root *TreeNode) []*TreeNode {
	var result, storey []*TreeNode
	if root == nil {
		return result
	}
	storey = append(result, root)

	hasNextStorey := true
	for hasNextStorey {
		curStorey := storey
		storey = nil
		hasNextStorey = false
		for _, it := range curStorey {
			result = append(result, it)
			if it == nil {
				storey = append(storey, nil, nil)
			} else {
				storey = append(storey, it.Left, it.Right)
				if it.Left != nil || it.Right != nil {
					hasNextStorey = true
				}
			}
		}
	}

	// Trim End
	count := len(result)
	for i := count - 1; i > 0; i-- {
		if result[i] != nil {
			break
		}
		count--
	}
	return result[:count]
}

const nullTreeValue = 0xffffffff
const null = nullTreeValue

func sliceToTree(arr []int) *TreeNode {
	var slice []*TreeNode

	for _, it := range arr {
		if it == nullTreeValue {
			slice = append(slice, nil)
		} else {
			slice = append(slice, &TreeNode{Val: it})
		}
	}

	size := len(slice)
	if size == 0 {
		return nil
	}
	for i := 0; i < size; i++ {
		it := slice[i]
		if it == nil {
			continue
		}
		j, k := (i<<1)+1, (i+1)<<1
		if j < size {
			it.Left = slice[j]
		} else {
			break
		}
		if k < size {
			it.Right = slice[k]
		} else {
			break
		}
	}

	return slice[0]
}
