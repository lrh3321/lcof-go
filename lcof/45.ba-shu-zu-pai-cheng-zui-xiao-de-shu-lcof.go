// [面试题45. 把数组排成最小的数](https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/)

package main

import (
	"bytes"
	"sort"
	"strconv"
)

type stringSlice []string

func (p stringSlice) Len() int           { return len(p) }
func (p stringSlice) Less(i, j int) bool { return p[i]+p[j] < p[j]+p[i] }
func (p stringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func minNumber(nums []int) string {
	var strs = make([]string, len(nums), len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	sort.Sort(stringSlice(strs))
	w := bytes.NewBuffer(nil)
	for _, s := range strs {
		w.WriteString(s)
	}
	return w.String()
}

func minNumberMain() {
	println(minNumber([]int{10, 2}), "102")
	println(minNumber([]int{3, 30, 34, 5, 9}), "3033459")
}
