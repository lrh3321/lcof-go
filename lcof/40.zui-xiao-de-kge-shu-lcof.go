// [面试题40. 最小的k个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)

package main

import (
	"container/heap"
	"fmt"
)

type intHeap []int

func (p intHeap) Len() int           { return len(p) }
func (p intHeap) Less(i, j int) bool { return p[i] > p[j] }
func (p intHeap) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *intHeap) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func (p *intHeap) Pop() interface{} {
	size := len(*p)
	v := (*p)[size-1]
	*p = (*p)[:size-1]
	return v
}

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	size := len(arr)
	if size == k {
		return arr
	}
	heap1 := intHeap(arr[:k])

	heap.Init(&heap1)

	for _, it := range arr[k:] {
		if it < heap1[0] {
			heap.Push(&heap1, it)
			heap.Pop(&heap1)
		}
	}

	return heap1
}

func getLeastNumbersMain() {
	fmt.Println(getLeastNumbers([]int{0, 1, 1, 2, 4, 4, 1, 3, 3, 2}, 6), []int{0, 1, 1, 1, 2, 2})

	fmt.Println(getLeastNumbers([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4), []int{1, 2, 3, 4})

	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2), []int{1, 2})
	fmt.Println(getLeastNumbers([]int{0, 1, 2, 1}, 1), []int{0})
	fmt.Println(getLeastNumbers([]int{0, 0, 2, 3, 2, 1, 1, 2, 0, 4}, 10), []int{0, 0, 2, 3, 2, 1, 1, 2, 0, 4})
	fmt.Println(getLeastNumbers([]int{0, 0, 0, 2, 0, 5}, 0), []int{})

}
