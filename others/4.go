// +build ignore

package main

import "fmt"

// 帕斯卡三角形
// 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

// 在杨辉三角中，每个数是它左上方和右上方的数的和。

// 示例:

// 输入: 5
// 输出:
// [
//      [1],
//     [1,1],
//    [1,2,1],
//   [1,3,3,1],
//  [1,4,6,4,1]
// ]

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	ret := [][]int{{1}}
	before := []int{1}
	for i := 1; i < numRows; i++ {
		tmp := make([]int, len(before)+1)
		tmp[0] = 1
		tmp[len(before)] = 1
		for j := 1; j < len(before); j++ {
			tmp[j] = before[j-1] + before[j]
		}
		ret = append(ret, tmp[0:])
		before = tmp[0:]
	}
	return ret
}

func main() {
	fmt.Println(generate(0))
}
