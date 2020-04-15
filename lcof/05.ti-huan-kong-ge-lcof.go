// [面试题05. 替换空格](https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/)
package main

import "bytes"

func replaceSpace(s string) string {
	w := bytes.NewBuffer(nil)
	space := []byte{'%', '2', '0'}
	for _, c := range s {
		if c == ' ' {
			w.Write(space)
		} else {
			w.WriteRune(c)
		}
	}
	return w.String()
}

func replaceSpaceMain() {
	println(replaceSpace("We are happy."))
}
