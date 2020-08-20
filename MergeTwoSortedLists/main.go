package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		if l1 != nil && (l2 == nil || l1.Val <= l2.Val) {
			cur.Next = &ListNode{
				Val:  l1.Val,
				Next: nil,
			}
			l1 = l1.Next
			cur = cur.Next
		} else if l2 != nil && (l1 == nil || l2.Val <= l1.Val) {
			cur.Next = &ListNode{
				Val:  l2.Val,
				Next: nil,
			}
			l2 = l2.Next
			cur = cur.Next
		}
	}
	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{}
	fmt.Println(head.Next)
}
