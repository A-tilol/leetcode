package main

// ListNode is definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var tail *ListNode
	cur := head
	for cur != nil {
		tail = &ListNode{cur.Val, tail}
		cur = cur.Next
	}

	for head != nil {
		if head.Val != tail.Val {
			return false
		}
		head, tail = head.Next, tail.Next
	}

	return true
}

func main() {
}
