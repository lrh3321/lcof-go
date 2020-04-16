// [面试题52. 两个链表的第一个公共节点](https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/)

package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != nil && curB != nil {
		curA = curA.Next
		curB = curB.Next
	}

	if curA != nil {
		var k int
		for curA != nil {
			curA = curA.Next
			k++
		}
		curA, curB = headA, headB
		for i := 0; i < k; i++ {
			curA = curA.Next
		}
	} else if curB != nil {
		var k int
		for curB != nil {
			curB = curB.Next
			k++
		}
		curA, curB = headA, headB
		for i := 0; i < k; i++ {
			curB = curB.Next
		}
	} else {
		curA, curB = headA, headB
	}

	for curA != nil && curB != nil {
		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}

	return nil
}

func getIntersectionNodeMain() {
	var headA, headB, headC, headD *ListNode

	headA = sliceToListNode([]int{4, 1, 8, 4, 5})
	headC = headA
	for i := 0; i < 2; i++ {
		headC = headC.Next
	}
	headB = sliceToListNode([]int{5, 0, 1})
	headD = headB
	for headD.Next != nil {
		headD = headD.Next
	}
	headD.Next = headC
	fmt.Println(
		getIntersectionNode(
			headA,
			headB,
		),
		headC,
	)

	headA = sliceToListNode([]int{0, 9, 1, 2, 4})
	headC = headA
	for i := 0; i < 3; i++ {
		headC = headC.Next
	}
	headB = sliceToListNode([]int{3})
	headD = headB
	for headD.Next != nil {
		headD = headD.Next
	}
	headD.Next = headC
	fmt.Println(
		getIntersectionNode(
			headA,
			headB,
		),
		headC,
	)

	fmt.Println(
		getIntersectionNode(
			sliceToListNode([]int{2, 6, 4}),
			sliceToListNode([]int{1, 5}),
		),
		nil,
	)
}
