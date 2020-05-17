// +build ignore

package main

import "fmt"

/*

二叉树的层序遍历
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。



示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return [][]int{}
	}

	level := []*TreeNode{}
	level = append(level, root)

	for {
		tmp := []int{}
		tmpLevel := []*TreeNode{}
		for _, elem := range level {
			if elem == nil {
				continue
			}
			tmp = append(tmp, elem.Val)
			tmpLevel = append(tmpLevel, elem.Left, elem.Right)
		}
		if len(tmpLevel) == 0 {
			break
		}
		level = tmpLevel
		ret = append(ret, tmp)
	}
	return ret
}

func main() {
	root := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{
		Val: 1,
		// Left:  &TreeNode{},
		// Right: &TreeNode{},
	}

	fmt.Println(levelOrder(root))

}
