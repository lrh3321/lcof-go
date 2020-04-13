// [面试题06. 从尾到头打印链表](https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/)

package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	res := reversePrint(head.Next)

	return append(res, head.Val)
}

func reversePrintMain() {
	fmt.Println(reversePrint(sliceToListNode([]int{1, 3, 2})))
}
