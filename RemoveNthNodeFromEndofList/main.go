package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// less space complexity
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	size := 0
	for cur != nil {
		cur = cur.Next
		size++
	}

	if size == n {
		return head.Next
	}

	cur = head
	for i := 0; i < size-n-1; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next

	return head
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	nodes := make([]*ListNode, 0)
	cur := head
	size := 0
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
		size++
	}
	if size == n {
		return head.Next
	}
	nodes[size-n-1].Next = nodes[size-n-1].Next.Next
	return head
}

func main() {
}
