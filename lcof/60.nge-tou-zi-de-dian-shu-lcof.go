// [面试题60. n个骰子的点数](https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/)

package main

import "fmt"

func twoSum(n int) []float64 {
	factor := 1.0 / 6.0
	var memo = make([][]float64, n+1, n+1)
	memo[1] = []float64{0, factor, factor, factor, factor, factor, factor}
	for i := 2; i <= n; i++ {
		min, max := i, 6*i
		row := make([]float64, max+1)
		left, right := min, max
		memo[i] = row
		for left <= right {
			p := 0.0
			for j := 1; j <= 6; j++ {
				k := left - j
				if k <= 0 {
					break
				}
				p += memo[i-1][k] * factor
			}
			row[left], row[right] = p, p
			left++
			right--
		}
	}

	return memo[n][n:]
}

func twoSumMain() {
	fmt.Println(twoSum(1), []float64{0.16667, 0.16667, 0.16667, 0.16667, 0.16667, 0.16667})
	fmt.Println(twoSum(2), []float64{0.02778, 0.05556, 0.08333, 0.11111, 0.13889, 0.16667, 0.13889, 0.11111, 0.08333, 0.05556, 0.02778})
}
