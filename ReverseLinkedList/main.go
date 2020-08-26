package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var ans, pre *ListNode = nil, nil
	for head != nil {
		ans = &ListNode{head.Val, pre}
		pre = ans
		head = head.Next
	}
	return ans
}

func main() {
}
