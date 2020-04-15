// [面试题65. 不用加减乘除做加法](https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/)

package main

func add(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	for b != 0 {
		plus := (a ^ b)    // 求和（不计进位）. 相同位置0，相反位置1
		b = ((a & b) << 1) // 计算进位. 先保留同为1的位，都为1的位要向左进位，因此左移1位
		a = plus
	}
	return a
}

func addMain() {
	var a, b int
	a, b = 1, 1
	println(add(a, b), a+b)

	a, b = 11, 19
	println(add(a, b), a+b)
}
