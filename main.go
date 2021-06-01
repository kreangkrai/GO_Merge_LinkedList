package main

import (
	"fmt"
	"sort"
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

func mergeKLists(lists []*ListNode) *ListNode {
	n := []int{}
	if len(lists) == 0 {
		return nil
	}
	for i := 0; i < len(lists); i++ {
		tmp := lists[i]
		for tmp != nil {
			n = append(n, tmp.Val)
			tmp = tmp.Next
		}
	}
	if len(n) == 0 {
		return nil
	}
	sort.Ints(n)
	start := &ListNode{Val: n[0]}
	tmp1 := start

	for i := 0; i < len(n); i++ {
		tmp1.Next = &ListNode{Val: n[i]}
		tmp1 = tmp1.Next
	}
	return start.Next

}
func main() {
	l := []*ListNode{}
	l1 := &ListNode{Val: 1}
	l1.AddNode(4)
	l1.AddNode(5)
	l = append(l, l1)
	l2 := &ListNode{Val: 1}
	l1.AddNode(3)
	l2.AddNode(4)
	l = append(l, l2)
	l3 := &ListNode{Val: 2}
	l3.AddNode(15)
	l = append(l, l3)

	l4 := &ListNode{Val: 9}
	l = append(l, l4)
	k := mergeKLists(l)
	k.PrintNode()

}
