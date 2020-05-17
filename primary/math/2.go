// +build ignore

package main

import "fmt"

/*
计数质数
统计所有小于非负整数 n 的质数的数量。

示例:

输入: 10
输出: 4
解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
*/

func countPrimes(n int) int {
	// 2,3,4,5,6,7,8
	ret := 0
	if n < 2 {
		return ret
	}
	mp := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if !mp[i] {
			for j := i * i; j < n; j = j + i {
				mp[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if !mp[i] {
			ret++
		}
	}
	return ret
}

func main() {
	fmt.Println(countPrimes(3))
}
