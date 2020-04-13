// [面试题24. 反转链表](https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/)

package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var stack []*ListNode
	cur := head
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Next
	}
	if len(stack) > 0 {
		i := len(stack) - 1
		head = stack[i]
		cur = head
		for i > 0 {
			i--
			cur.Next = stack[i]
			cur = cur.Next
		}
		cur.Next = nil
	}

	return head
}

func reverseListMain() {
	fmt.Println(
		listNodeToSlice(reverseList(sliceToListNode([]int{1, 2, 3, 4, 5}))),
		[]int{5, 4, 3, 2, 1},
	)
}
