// [面试题29. 顺时针打印矩阵](https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/)

package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	if m == 1 {
		return matrix[0]
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}

	count := m * n
	var result = make([]int, count, count)
	copy(result, matrix[0])
	i := n
	t, b, l, r := 1, m-1, 0, n-1

	dir := 0
	for i < count {
		switch dir {
		case 0:
			for j := t; j <= b; j++ {
				result[i] = matrix[j][r]
				i++
			}
			r--
		case 1:
			row := matrix[b]
			for j := r; j >= l; j-- {
				result[i] = row[j]
				i++
			}
			b--
		case 2:
			for j := b; j >= t; j-- {
				result[i] = matrix[j][l]
				i++
			}
			l++
		case 3:
			row := matrix[t][l : r+1]
			copy(result[i:], row)
			i += len(row)
			t++
		}
		dir++
		dir %= 4
	}

	return result
}

func spiralOrderMain() {
	var matrix [][]int
	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(spiralOrder(matrix))
	fmt.Println([]int{1, 2, 3, 6, 9, 8, 7, 4, 5})

	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Println(spiralOrder(matrix))
	fmt.Println([]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7})
}
