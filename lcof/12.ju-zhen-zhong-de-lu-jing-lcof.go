// [面试题12. 矩阵中的路径](https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/)

package main

func existDFS(board [][]byte, visited [][]bool, target []byte, row, col, m, n int) bool {
	if len(target) == 0 {
		return true
	}
	first := target[0]
	target = target[1:]

	contains := false
	if row > 0 && board[row-1][col] == first && !visited[row-1][col] && !contains {
		visited[row-1][col] = true
		contains = existDFS(board, visited, target, row-1, col, m, n)

		visited[row-1][col] = false
	}
	if row+1 < m && board[row+1][col] == first && !visited[row+1][col] && !contains {
		visited[row+1][col] = true
		contains = existDFS(board, visited, target, row+1, col, m, n)

		visited[row+1][col] = false
	}
	if col > 0 && board[row][col-1] == first && !visited[row][col-1] && !contains {
		visited[row][col-1] = true
		contains = existDFS(board, visited, target, row, col-1, m, n)

		visited[row][col-1] = false
	}
	if col+1 < n && board[row][col+1] == first && !visited[row][col+1] && !contains {
		visited[row][col+1] = true
		contains = existDFS(board, visited, target, row, col+1, m, n)

		visited[row][col+1] = false
	}

	return contains
}

func exist(board [][]byte, word string) bool {

	target := []byte(word)
	if len(target) == 0 {
		return true
	}
	m, n := len(board), len(board[0])
	visited := make([][]bool, m, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n, n)
	}

	first := target[0]
	target = target[1:]
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if board[row][col] != first {
				continue
			}
			visited[row][col] = true
			if existDFS(board, visited, target, row, col, m, n) {
				return true
			}
			visited[row][col] = false
		}
	}
	return false
}
func existMain() {
	var board [][]byte
	var word string
	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word = "ABCCED"
	println(exist(board, word), true)

	board = [][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}
	word = "abcd"
	println(exist(board, word), false)
}
