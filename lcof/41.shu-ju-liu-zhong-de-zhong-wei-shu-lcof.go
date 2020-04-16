// [面试题41. 数据流中的中位数](https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/)

package main

import (
	"container/heap"
)

type maxHeap []int

func (a maxHeap) Len() int           { return len(a) }
func (a maxHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a maxHeap) Less(i, j int) bool { return a[i] > a[j] }
func (a *maxHeap) Push(x interface{}) {
	*a = append(*a, x.(int))
}
func (a *maxHeap) Pop() interface{} {
	i := len(*a) - 1
	v := (*a)[i]
	*a = (*a)[:i]
	return v
}

type minHeap []int

func (a minHeap) Len() int           { return len(a) }
func (a minHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a minHeap) Less(i, j int) bool { return a[i] < a[j] }
func (a *minHeap) Push(x interface{}) {
	*a = append(*a, x.(int))
}
func (a *minHeap) Pop() interface{} {
	i := len(*a) - 1
	v := (*a)[i]
	*a = (*a)[:i]
	return v
}

type MedianFinder struct {
	left  *maxHeap // 小于等于中位数
	right *minHeap // 大于等于中位数
	count int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		left:  &maxHeap{},
		right: &minHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {

	if len(*this.left) == 0 || num < (*this.left)[0] {
		heap.Push(this.left, num)
		if this.left.Len()-this.right.Len() > 1 {
			v := heap.Pop(this.left).(int)
			// println("Top:", v)
			heap.Push(this.right, v)
		}
	} else {
		heap.Push(this.right, num)
		if this.right.Len()-this.left.Len() > 1 {
			v := heap.Pop(this.right).(int)
			// println("Top:", v)
			heap.Push(this.left, v)
		}
	}

	// fmt.Println("Left:", this.left)
	// fmt.Println("Right:", this.right)
	this.count++
}

func (this *MedianFinder) FindMedian() float64 {
	if this.count&1 == 0 {
		return float64((*this.left)[0]+(*this.right)[0]) / 2.0
	}
	if this.right.Len() > this.left.Len() {
		return float64((*this.right)[0])
	}
	return float64((*this.left)[0])

}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func MedianFinderMain() {
	var obj MedianFinder

	obj = Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	println(obj.FindMedian(), 1.5)
	obj.AddNum(3)
	println(obj.FindMedian(), 2.0)

	obj = Constructor()
	obj.AddNum(-1)
	println(obj.FindMedian(), -1.0)
	obj.AddNum(-2)
	println(obj.FindMedian(), -1.5)
	obj.AddNum(-3)
	println(obj.FindMedian(), -2.0)
	obj.AddNum(-4)
	println(obj.FindMedian(), -2.5)
	obj.AddNum(-5)
	println(obj.FindMedian(), -3.0)

	obj = Constructor()
	obj.AddNum(6)
	println(obj.FindMedian(), 6.0)
	obj.AddNum(10)
	println(obj.FindMedian(), 8.0)
	obj.AddNum(2)
	println(obj.FindMedian(), 6.0)
	obj.AddNum(6)
	println(obj.FindMedian(), 6.0)
	obj.AddNum(5)
	println(obj.FindMedian(), 6.0)
	obj.AddNum(0)
	println(obj.FindMedian(), 5.5)
	obj.AddNum(6)
	println(obj.FindMedian(), 6.0)
	obj.AddNum(3)
	println(obj.FindMedian(), 5.5)
	obj.AddNum(1)
	println(obj.FindMedian(), 5.0)
	obj.AddNum(0)
	println(obj.FindMedian(), 4.0)
	obj.AddNum(0)
	println(obj.FindMedian(), 3.0)
}
