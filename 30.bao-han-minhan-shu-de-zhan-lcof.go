// [面试题30. 包含min函数的栈](https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/)

package main

type MinStack struct {
	s1, s2 []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.s1 = append(this.s1, x)
	size := len(this.s2)
	if size == 0 || this.s2[size-1] >= x {
		this.s2 = append(this.s2, x)
	}
}

func (this *MinStack) Pop() {
	v := this.Top()
	this.s1 = this.s1[:len(this.s1)-1]
	if v == this.s2[len(this.s2)-1] {
		this.s2 = this.s2[:len(this.s2)-1]
	}
}

func (this *MinStack) Top() int {
	return this.s1[len(this.s1)-1]
}

func (this *MinStack) Min() int {
	return this.s2[len(this.s2)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
