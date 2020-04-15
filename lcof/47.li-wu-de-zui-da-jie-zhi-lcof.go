// [面试题47. 礼物的最大价值](https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/)

package main

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m, m)

	sum := 0
	firstRow := make([]int, n, n)
	dp[0] = firstRow
	grid0 := grid[0]
	for i := 0; i < n; i++ {
		sum += grid0[i]
		firstRow[i] = sum
	}
	sum = grid0[0]
	for i := 1; i < m; i++ {
		dp[i] = make([]int, n, n)
		sum += grid[i][0]
		dp[i][0] = sum
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			var t1, t2 int
			t1, t2 = dp[row-1][col], dp[row][col-1]
			if t2 > t1 {
				t1 = t2
			}
			dp[row][col] = t1 + grid[row][col]
		}
	}
	return dp[m-1][n-1]
}

func maxValueMain() {
	println(maxValue([][]int{
		{1, 3, 1},
		{1, 5, 1},
	}), 10)
	println(maxValue([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}), 12)
}
