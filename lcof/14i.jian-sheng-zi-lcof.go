// [面试题14- I. 剪绳子](https://leetcode-cn.com/problems/jian-sheng-zi-lcof/)

package main

func cuttingRopeI(n int) int {
	dp := make([]int, n+1, n+1)
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		end := (i + 1) >> 1
		for j := 1; j <= end; j++ {
			k := i - j
			if dp[k] > k {
				k = dp[k]
			}
			t := j * k
			if t > dp[i] {
				dp[i] = t
			}
		}
	}

	// fmt.Println(dp)
	return dp[n]
}
