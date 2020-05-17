// +build ignore

package main

import "fmt"

/*
回文链表
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {

	// 空间复杂度不符合
	// s := []int{}

	// for head != nil {
	// 	s = append(s, head.Val)
	// 	head = head.Next
	// }
	// length := len(s)
	// for i := 0; i < length/2; i++ {
	// 	if s[i] != s[length-i-1] {
	// 		return false
	// 	}
	// }
	// return true

	// 翻转一半链表
	h := head
	l := 0
	for h != nil {
		h = h.Next
		l++
	}
	h = head
	for i := 0; i < (l+1)/2; i++ {
		h = h.Next
	}
	h = reverseList(h)
	for h != nil {
		if h.Val != head.Val {
			return false
		}
		h = h.Next
		head = head.Next
	}
	return true

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
		Val:  1,
		Next: nil,
	}
	root.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	root.Next.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	root.Next.Next.Next = &ListNode{
		Val:  1,
		Next: nil,
	}
	fmt.Println(isPalindrome(root))
}
