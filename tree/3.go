// +build ignore

package main

import "fmt"

/*
对称二叉树
给定一个二叉树，检查它是否是镜像对称的。



例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3


但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3


进阶：

你可以运用递归和迭代两种方法解决这个问题吗？
*/

/**
  Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetric2(root.Left, root.Right)
}

func isSymmetric2(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return isSymmetric2(left.Left, right.Right) && isSymmetric2(left.Right, right.Left)
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

	fmt.Println(isSymmetric(root))

}
