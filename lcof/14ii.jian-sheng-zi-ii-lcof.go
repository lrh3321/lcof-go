// [面试题14- II. 剪绳子 II](https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/)

package main

func cuttingRope(n int) int {
	switch n {
	case 2:
		return 1
	case 3:
		return 2
	case 4:
		return 4
	}

	var result = 1
	for n > 4 {
		result *= 3
		if result >= 1000000007 {
			result %= 1000000007
		}
		n -= 3
	}
	return (result * n) % 1000000007
}

func cuttingRopeMain() {
	println(cuttingRope(2), 1)
	println(cuttingRope(4), 4)
	println(cuttingRope(10), 36)
	println(cuttingRope(58), 549681949)

	println(cuttingRope(100), 703522804)
	println(cuttingRope(110), 492673366)
	println(cuttingRope(150), 710104287)
	println(cuttingRope(200), 818838607)
	println(cuttingRope(1000), 620946522)
}
