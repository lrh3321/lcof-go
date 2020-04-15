// [面试题20. 表示数值的字符串](https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/)

package main

import "strings"

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" || s == "." || s == "+" || s == "-" {
		return false
	}
	if len(s) >= 2 {
		first, second := s[0], s[1]
		if second == 'e' || second == 'E' {
			if first == '+' || first == '-' || first == '.' {
				return false
			}
		}
	}

	dot := false
	EIndex := -1
	// 有数字出现
	digit := false
	for i, c := range s {
		switch c {
		case '+', '-':
			if i != 0 {
				if EIndex+1 != i {
					return false
				}
			}
			digit = false
		case '.':
			if dot || EIndex > 0 {
				return false
			}
			dot = true
		case 'e', 'E':
			if i == 0 {
				return false
			}

			if EIndex > 0 {
				return false
			}
			EIndex = i
			digit = false
		default:
			if c > '9' || c < '0' {
				return false
			}
			digit = true
		}
	}
	return digit
}
func isNumberMain() {

	for _, s := range []string{"+100", "5e2", "-123", "3.1416", "0123", "-1E-16"} {
		println(isNumber(s), true, s)
	}
	for _, s := range []string{"", "e", "12e", "1a3.14", "1.2.3", "+-5", "12e+5.4", ".", "+", "-", ".e1", "+e1", "-e1", "4e+", "-."} {
		println(isNumber(s), false, s)
	}
}
