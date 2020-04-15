// [面试题22. 链表中倒数第k个节点](https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/)

package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	var bucket []*ListNode
	cur := head
	for cur != nil {
		bucket = append(bucket, cur)
		cur = cur.Next
	}

	if k-1 >= len(bucket) {
		return head
	}
	return bucket[len(bucket)-k]
}

func getKthFromEndMain() {
	fmt.Println(
		listNodeToSlice(getKthFromEnd(sliceToListNode([]int{1, 2, 3, 4, 5}), 2)),
		[]int{4, 5},
	)
	fmt.Println(
		listNodeToSlice(getKthFromEnd(sliceToListNode([]int{1, 2, 3, 4, 5}), 5)),
		[]int{1, 2, 3, 4, 5},
	)
	fmt.Println(
		listNodeToSlice(getKthFromEnd(sliceToListNode([]int{1, 2, 3, 4, 5}), 6)),
		[]int{1, 2, 3, 4, 5},
	)
}
