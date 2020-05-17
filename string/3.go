// +build ignore

package main

import "fmt"

/*
字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.


注意事项：您可以假定该字符串只包含小写字母。
*/

type Elem struct {
	index int
	cout  int
}

func firstUniqChar(s string) int {
	m := [26]Elem{}
	for i, c := range s {
		m[c-'a'].index = i
		m[c-'a'].cout++
	}
	minIndex := -1
	for _, elem := range m {
		if elem.cout == 1 {
			if minIndex == -1 {
				minIndex = elem.index
				continue
			}
			if elem.index < minIndex {
				minIndex = elem.index
			}
		}
	}
	return minIndex
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}
