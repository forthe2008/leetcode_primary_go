package main

import "fmt"

/*
 删除链表中的节点
请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。

现有一个链表 -- head = [4,5,1,9]，它可以表示为:





示例 1:

输入: head = [4,5,1,9], node = 5
输出: [4,1,9]
解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
示例 2:

输入: head = [4,5,1,9], node = 1
输出: [4,5,9]
解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.


说明:

链表至少包含两个节点。
链表中所有节点的值都是唯一的。
给定的节点为非末尾节点并且一定是链表中的一个有效节点。
不要从你的函数中返回任何结果。
*/

//ListNode  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	root := &ListNode{
		Val:  4,
		Next: nil,
	}
	root.Next = &ListNode{
		Val:  5,
		Next: nil,
	}
	root.Next.Next = &ListNode{
		Val:  1,
		Next: nil,
	}
	root.Next.Next.Next = &ListNode{
		Val:  9,
		Next: nil,
	}
	deleteNode(root.Next)
	fmt.Println(root)
	fmt.Println(root.Next)
	fmt.Println(root.Next.Next)
}
