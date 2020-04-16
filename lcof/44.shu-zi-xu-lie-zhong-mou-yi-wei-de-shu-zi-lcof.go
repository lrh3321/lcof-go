// [面试题44. 数字序列中某一位的数字](https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/)

package main

func findNthDigit(n int) int {
	base, digits := 9, 1

	for n > (base * digits) {
		n -= base * digits
		base *= 10
		digits++
	}
	// println("n:", n)

	i := n % digits
	if i == 0 {
		i = digits
	}

	pow := 1
	for j := 1; j < digits; j++ {
		pow *= 10
	}
	// println("firstNum:", pow)
	pow += n / digits

	if i == digits {
		pow--
	}

	for j := i; j < digits; j++ {
		pow /= 10
	}

	return pow % 10
}

func findNthDigitMain() {
	println(findNthDigit(3), 3)
	println(findNthDigit(11), 0)

	println(findNthDigit(10), 1)

	println(findNthDigit(99), 4)
	println(findNthDigit(100), 5)
	println(findNthDigit(101), 5)
	println(findNthDigit(10000), 7)
}
