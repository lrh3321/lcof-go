// [面试题13. 机器人的运动范围](https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/)

package main

func movingCount(m int, n int, k int) int {
	var result = 1
	if k == 0 {
		return result
	}
	visited := make([][]bool, m, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n, n)
	}
	visited[0][0] = true
	nextSlice := []int{0, 0}

	var canEnter = func(row, col int) bool {
		visited[row][col] = true
		var val int
		for row > 0 {
			val += row % 10
			row /= 10
		}
		for col > 0 {
			val += col % 10
			col /= 10
		}
		return val <= k
	}

	for len(nextSlice) > 0 {
		slice := nextSlice
		sliceSize := len(slice)
		nextSlice = nil

		for i := 0; i < sliceSize; i += 2 {
			row, col := slice[i], slice[i+1]
			if row > 0 && !visited[row-1][col] && canEnter(row-1, col) {
				result++
				visited[row-1][col] = true
				nextSlice = append(nextSlice, row-1, col)
			}
			if row+1 < m && !visited[row+1][col] && canEnter(row+1, col) {
				result++
				visited[row+1][col] = true
				nextSlice = append(nextSlice, row+1, col)
			}
			if col > 0 && !visited[row][col-1] && canEnter(row, col-1) {
				result++
				visited[row][col-1] = true
				nextSlice = append(nextSlice, row, col-1)
			}
			if col+1 < n && !visited[row][col+1] && canEnter(row, col+1) {
				result++
				visited[row][col+1] = true
				nextSlice = append(nextSlice, row, col+1)
			}
		}
	}

	return result
}

func movingCountMain() {
	println(movingCount(2, 3, 1), 3)
	println(movingCount(3, 1, 0), 1)
}
