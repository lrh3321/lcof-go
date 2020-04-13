// [面试题10- I. 斐波那契数列](https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/)

package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var mem = make([]int, n+1, n+1)
	mem[1] = 1
	for i := 2; i <= n; i++ {
		t := mem[i-1] + mem[i-2]
		if t >= 1000000007 {
			t %= 1000000007
		}
		mem[i] = t
	}
	return mem[n]
}

func fibMain() {
	fmt.Println(fib(2), 1)
	fmt.Println(fib(5), 5)
}
