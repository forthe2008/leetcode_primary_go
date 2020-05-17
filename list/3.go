// +build ignore

package main

import "fmt"

/*
反转链表
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
*/

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	fast := head
	slow := head
	if head == nil {
		return nil
	}
	for fast.Next != nil {
		fast = fast.Next
	}

	for slow != fast {
		tmp := slow.Next
		slow.Next = fast.Next
		fast.Next = slow
		slow = tmp
	}
	return fast
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

	r := reverseList(root)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}

}
