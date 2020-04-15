// [面试题10- II. 青蛙跳台阶问题](https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/)

package main

func numWays(n int) int {
	if n < 2 {
		return 1
	}
	var mem = make([]int, n+1, n+1)
	mem[0] = 1
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
func numWaysMain() {
	println(numWays(2), 2)
	println(numWays(7), 21)
}
