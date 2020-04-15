// [面试题25. 合并两个排序的链表](https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/)

package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head *ListNode
	if l2.Val > l1.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	cur := head

	for l1 != nil && l2 != nil {
		if l2.Val > l1.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 == nil {
		cur.Next = l2
	}

	if l2 == nil {
		cur.Next = l1
	}

	return head
}

func mergeTwoListsMain() {
	fmt.Println(
		listNodeToSlice(mergeTwoLists(
			sliceToListNode([]int{1, 2, 4}),
			sliceToListNode([]int{1, 3, 4}),
		)),
		[]int{1, 1, 2, 3, 4, 4},
	)
}
