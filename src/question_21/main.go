package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l0 *ListNode = new(ListNode)
	var l1 *ListNode = new(ListNode)
	var l2 *ListNode = new(ListNode)
	var l3 *ListNode = new(ListNode)
	var l4 *ListNode = new(ListNode)

	l0.Next = l1
	l1.Val = 1
	l1.Next = l2
	l2.Val = 2
	l2.Next = l3
	l3.Val = 4
	l3.Next = l4
	l4.Val = 8

	temp := l0.Next
	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var current *ListNode = new(ListNode)
	head = current
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}
	return head.Next
}
