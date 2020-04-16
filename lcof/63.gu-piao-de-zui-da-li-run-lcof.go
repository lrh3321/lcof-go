// [面试题63. 股票的最大利润](https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/)

package main

func maxProfit(prices []int) int {
	var ans int
	size := len(prices)
	if size == 0 {
		return 0
	}
	low := prices[0]

	for i := 0; i < size; i++ {
		it := prices[i]
		if it > low {
			t := it - low
			if t > ans {
				ans = t
			}
			continue
		}
		low = it
	}

	return ans
}

func maxProfitMain() {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}), 5)
	println(maxProfit([]int{7, 6, 4, 3, 1}), 0)
}
