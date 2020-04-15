// [面试题18. 删除链表的节点](https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/)

package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	head.Next = deleteNode(head.Next, val)
	return head
}

func deleteNodeMain() {

	fmt.Println(listNodeToSlice(deleteNode(sliceToListNode([]int{4, 5, 1, 9}), 5)))
	fmt.Println(listNodeToSlice(deleteNode(sliceToListNode([]int{4, 5, 1, 9}), 1)))

	fmt.Println(listNodeToSlice(deleteNode(sliceToListNode([]int{4, 5, 1, 9}), 10)))
}
