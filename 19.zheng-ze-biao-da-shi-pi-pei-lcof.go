// [面试题19. 正则表达式匹配](https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof/)

package main

func isMatch(s string, p string) bool {
	ns, np := len(s), len(p)
	dp := make([][]bool, ns+1, ns+1)
	for i := 0; i <= ns; i++ {
		dp[i] = make([]bool, np+1, np+1)
	}

	for i := 0; i <= ns; i++ {
		dp[i][0] = i == 0
		for j := 1; j <= np; j++ {
			pChar := p[j-1]
			if pChar == '*' {
				if j >= 2 {
					dp[i][j] = dp[i][j-2]

					if i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') {
						dp[i][j] = dp[i][j] || dp[i-1][j]
					}
				}
			} else {
				if i > 0 && (s[i-1] == pChar || pChar == '.') {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}

	return dp[ns][np]
}

func isMatchMain() {
	println(isMatch("aa", "a"), false)
	println(isMatch("aa", "a*"), true)
	println(isMatch("ab", ".*"), true)
	println(isMatch("aab", "c*a*b"), true)
	println(isMatch("mississippi", "mis*is*p*."), false)
}
