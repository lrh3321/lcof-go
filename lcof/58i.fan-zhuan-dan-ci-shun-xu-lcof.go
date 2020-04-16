// [面试题58 - I. 翻转单词顺序](https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/)

package main

import (
	"bytes"
	"strings"
)

func reverseWords(s string) string {
	seq := strings.Split(s, " ")
	w := bytes.NewBuffer(nil)
	for i := len(seq) - 1; i >= 0; i-- {
		it := seq[i]
		if it == "" {
			continue
		}
		w.WriteString(it)
		w.WriteByte(' ')
	}
	str := w.String()
	if str != "" {
		str = str[:len(str)-1]
	}
	return str
}

func reverseWordsMain() {
	println(reverseWords("  hello world!  "), "world! hello")
	println(reverseWords("a good   example"), "example good a")
}
