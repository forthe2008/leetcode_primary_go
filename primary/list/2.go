// +build ignore

package main

import "fmt"

/*
删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head
	var pre *ListNode
	for i := 0; i < n && fast != nil; i++ {
		fast = fast.Next
	}
	for fast != nil {
		pre = slow
		fast = fast.Next
		slow = slow.Next
	}

	if slow.Next != nil {
		slow.Val = slow.Next.Val
		slow.Next = slow.Next.Next
	} else {
		if pre == nil {
			return nil
		}
		pre.Next = slow.Next
	}

	return head
}

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	removeNthFromEnd(head, 1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}
