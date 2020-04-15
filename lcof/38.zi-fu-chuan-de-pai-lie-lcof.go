// [面试题38. 字符串的排列](https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/)

package main

import "fmt"

func permutationImpl(needs map[rune]int, runeList []rune, remain int, prefix []rune, prefixLen int, result *[]string) {
	if remain == 0 {
		*result = append(*result, string(prefix))
		return
	}

	for _, c := range runeList {
		need := needs[c]
		if need == 0 {
			continue
		}

		needs[c]--
		prefix[prefixLen] = c
		permutationImpl(needs, runeList, remain-1, prefix, prefixLen+1, result)
		needs[c]++
	}
}
func permutation(s string) []string {
	var runeList []rune
	var needs = make(map[rune]int)
	var remain = len(s)
	for _, c := range s {
		if _, ok := needs[c]; !ok {
			runeList = append(runeList, c)
		}
		needs[c]++
	}

	var result []string
	permutationImpl(needs, runeList, remain, make([]rune, remain), 0, &result)
	return result
}

func permutationMain() {
	fmt.Println(permutation("abc"))
}
