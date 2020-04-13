// [面试题16. 数值的整数次方](https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/)

package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	negative := n < 0
	if negative {
		n = -n
	}

	var result = x
	for n > 1 {
		result = result * x
		if n%2 == 1 {
			result = result * x
		}
		x = x * x

		n >>= 1
	}

	if negative {
		return 1.0 / result
	}

	return result
}
func myPowMain() {
	println(myPow(2.00000, 5), 32.00000)
	println(myPow(2.00000, 10), 1024.00000)
	println(myPow(2.10000, 3), 9.26100)
	println(myPow(2.00000, -2), 0.25000)
}
