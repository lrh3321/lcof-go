// [面试题04. 二维数组中的查找](https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/)

package main

func findNumberIn2DArrayImpl(matrix [][]int, target, x0, y0, x1, y1 int) bool {
	// println(x0, y0, x1, y1)

	min, max := matrix[x0][y0], matrix[x1][y1]
	// println(min, max)
	if target == min || target == max {
		return true
	} else if target < min || target > max {
		return false
	}

	if x0 == x1 {
		if y0 == y1 {
			return false
		}
		arr := matrix[x0]
		y := (y0 + y1) >> 1
		for y > y0 {
			it := arr[y]
			if it == target {
				return true
			} else if it > target {
				y1 = y - 1
			} else {
				y0 = y + 1
			}
			y = (y0 + y1) >> 1
		}
		return arr[y0] == target || arr[y1] == target
	} else if y0 == y1 {
		x := (x0 + x1) >> 1
		for x > x0 {
			it := matrix[x][y0]
			if it == target {
				return true
			} else if it > target {
				x1 = x - 1
			} else {
				x0 = x + 1
			}
			x = (x0 + x1) >> 1
		}
		return matrix[x0][y0] == target || matrix[x1][y0] == target
	}

	x, y := (x0+x1)>>1, (y0+y1)>>1
	if x0 > x1 || y0 > y1 {
		return false
	}
	it := matrix[x][y]

	if x == x0 && y == y0 {
		// 2*2/1*2/2*1 的矩阵查找
		return matrix[x1][y0] == target || matrix[x0][y1] == target
	}
	// println(it)
	if it == target {
		return true
	} else if it > target {
		if y == y0 {
			return findNumberIn2DArrayImpl(matrix, target, x0, y0, x1, y0) ||
				findNumberIn2DArrayImpl(matrix, target, x0, y1, x1, y1)
		}
		if x == x0 {
			return findNumberIn2DArrayImpl(matrix, target, x0, y0, x0, y1) ||
				findNumberIn2DArrayImpl(matrix, target, x1, y0, x1, y1)
		}

		y2 := y - 1
		return findNumberIn2DArrayImpl(matrix, target, x0, y0, x1, y2) ||
			findNumberIn2DArrayImpl(matrix, target, x0, y, x-1, y1)
	}

	y2 := y + 1
	return findNumberIn2DArrayImpl(matrix, target, x0, y2, x1, y1) ||
		findNumberIn2DArrayImpl(matrix, target, x+1, y0, x1, y)

}
func findNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}

	x0, y0 := 0, 0

	return findNumberIn2DArrayImpl(matrix, target, x0, y0, n-1, m-1)
}

func findNumberIn2DArrayMain() {
	var matrix [][]int

	matrix = [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	println(findNumberIn2DArray(matrix, 5), true)

	matrix = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	println(findNumberIn2DArray(matrix, 5), true)
	println(findNumberIn2DArray(matrix, 20), false)
	println(findNumberIn2DArray(matrix, 0), false)
	println(findNumberIn2DArray(matrix, 31), false)
	println(findNumberIn2DArray(matrix[1:], 1), false)
	println(findNumberIn2DArray(matrix[1:], 31), false)

	matrix = [][]int{
		{-1, 3},
	}
	println(findNumberIn2DArray(matrix, 3), true)

	matrix = [][]int{
		{1, 3, 5},
	}
	println(findNumberIn2DArray(matrix, 1), true)

	matrix = [][]int{
		{1, 2, 3, 4, 5},
	}
	println(findNumberIn2DArray(matrix, 2), true)

	matrix = [][]int{
		{1, 4},
		{2, 5},
	}
	println(findNumberIn2DArray(matrix, 2), true)
}
