// [面试题09. 用两个栈实现队列](https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/)

package main

type stack struct {
	data []int
}

func (s *stack) Empty() bool {
	return len(s.data) == 0
}

func (s *stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *stack) Pop() int {
	val := s.data[0]
	s.data = s.data[1:]
	return val
}

type CQueue struct {
	s1, s2 *stack
}

func Constructor() CQueue {
	return CQueue{
		s1: &stack{},
		s2: &stack{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.s1.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.s2.Empty() {
		s1, s2 := this.s1, this.s2
		for !s1.Empty() {
			s2.Push(s1.Pop())
		}
	}
	if this.s2.Empty() {
		return -1
	}
	return this.s2.Pop()
}
