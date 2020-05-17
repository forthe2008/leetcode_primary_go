// +build ignore

package main

import (
	"fmt"
	"strings"
)

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false
*/

func isPalindrome(s string) bool {
	length := len(s)

	s = strings.ToLower(s)
	fmt.Println(s)
	for i, j := 0, length-1; i < j; {
		if (s[i] < 'a' || s[i] > 'z') && (s[i] < '0' || s[i] > '9') {
			i++
			continue
		}
		if (s[j] < 'a' || s[j] > 'z') && (s[j] < '0' || s[j] > '9') {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
