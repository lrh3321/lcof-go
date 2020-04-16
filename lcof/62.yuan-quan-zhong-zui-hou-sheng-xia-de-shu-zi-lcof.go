// [面试题62. 圆圈中最后剩下的数字](https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/)

package main

func lastRemaining(n int, m int) int {
	// 约瑟夫环
	var ans int
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}

func lastRemainingMain() {
	println(lastRemaining(5, 3), 3)
	println(lastRemaining(10, 17), 2)
}
