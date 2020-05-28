package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
func main() {
	var l10 *ListNode = new(ListNode)
	var l11 *ListNode = new(ListNode)
	var l12 *ListNode = new(ListNode)

	var l20 *ListNode = new(ListNode)
	var l21 *ListNode = new(ListNode)
	var l22 *ListNode = new(ListNode)


	l10.Val = 5
	l11.Val = 4
	l12.Val = 6
	l10.Next = l11
	l11.Next = l12

	l20.Val = 5
	l21.Val = 7
	l22.Val = 1
	l20.Next = l21
	l21.Next = l22

	res := addTwoNumbers(l10, l20)

	//test := revertNode(res)

	for res != nil{
		fmt.Println(res.Val)
		res = res.Next
	}
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode = new(ListNode)
	var head *ListNode
	head = result
	carry := 0
	for l1 != nil || l2 !=nil {
		d1:=0
		d2:=0
		if l1 == nil{//防止该节点为null
			d1 = 0
		}else{
			d1 = l1.Val
		}
		if l2 == nil{//防止该节点为null
			d2 = 0
		}else{
			d2 = l2.Val
		}

		sum := d1 + d2 + carry
		if sum >= 10 {
			carry = 1
		}else {
			carry = 0
		}
		var tmp *ListNode = new(ListNode)
		tmp.Val = sum % 10
		head.Next = tmp
		head = head.Next

		if l1 != nil{
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry == 1{
		var tmp *ListNode = new(ListNode)
		tmp.Val = 1
		head.Next = tmp
	}
	return result.Next
}
func revertNode(res *ListNode) *ListNode{

	var p *ListNode
	var q *ListNode
	var r *ListNode
	if res.Next == nil {
		fmt.Println(res.Val)

	}
	p = res
	q = res.Next
	res.Next = nil
	for q != nil {
		r = q.Next
		q.Next = p
		p = q
		q = r
	}
	res = p
	return res
}
