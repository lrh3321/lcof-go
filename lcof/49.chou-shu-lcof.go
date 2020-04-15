// [面试题49. 丑数](https://leetcode-cn.com/problems/chou-shu-lcof/)

package main

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n, n)
	dp[0] = 1

	var a, b, c int
	var min = func(i, j, k int) int {
		if j < i {
			i = j
		}
		if k < i {
			return k
		}
		return i
	}
	dpA, dpB, dpC := 2, 3, 5
	for i := 1; i < n; i++ {
		m := min(dpA, dpB, dpC)
		dp[i] = m
		if dpA == m {
			a++
			dpA = 2 * dp[a]
		}
		if dpB == m {
			b++
			dpB = 3 * dp[b]
		}
		if dpC == m {
			c++
			dpC = 5 * dp[c]
		}
	}

	return dp[n-1]
}

func nthUglyNumberMain() {
	println(nthUglyNumber(10), 12)

	println(nthUglyNumber(1690), 2123366400)
}
