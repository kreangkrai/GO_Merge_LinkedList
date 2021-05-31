package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) AddNode(val int) {
	newNode := ListNode{Val: val, Next: nil}
	iter := n
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = &newNode
}
func (n *ListNode) PrintNode() {
	iter := n
	for iter != nil {
		fmt.Print(iter.Val, " ")
		iter = iter.Next
	}
	fmt.Println()
}
func (n *ListNode) DeleteNode() {
	iter := n
	if iter.Next == nil {
		fmt.Println("Node is empty")
		return
	}
	for iter.Next.Next != nil {
		iter = iter.Next
	}
	iter.Next = nil
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var listNode = new(ListNode)
	var current = listNode
	for l1 != nil || l2 != nil {
		var value int
		if l1 == nil {
			value = l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			value = l1.Val
			l1 = l1.Next
		} else if l1.Val < l2.Val {
			value = l1.Val
			l1 = l1.Next
		} else {
			value = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	return listNode.Next
}
func main() {
	l1 := &ListNode{Val: 1}
	l1.AddNode(2)
	l1.AddNode(4)
	l1.PrintNode()
	l1.DeleteNode()
	l2 := &ListNode{Val: 1}
	l2.AddNode(3)
	l2.AddNode(4)
	l2.PrintNode()
	l2.DeleteNode()
	l2.DeleteNode()
	l2.DeleteNode()
	l2.DeleteNode()
	l2.AddNode(9)
	l2.PrintNode()

	k := mergeTwoLists(l1, l2)
	k.PrintNode()
}
