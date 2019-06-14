package addtwonumbers

import (
	"math/big"
	"strconv"
)

type ListNode struct {
	Val int

	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := ""
	cur := l1
	for {
		s1 = strconv.Itoa(cur.Val) + s1
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}
	s2 := ""
	cur = l2
	for {
		s2 = strconv.Itoa(cur.Val) + s2
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}

	a := new(big.Int)
	a, _ = a.SetString(s1, 10)
	b := new(big.Int)
	b, _ = b.SetString(s2, 10)
	s := a.Add(a, b).String()

	cur = nil
	for _, c := range s {
		v, _ := strconv.Atoi(string(c))
		cur = &ListNode{
			Val:  v,
			Next: cur,
		}
	}

	return cur
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1 := l1
	cur2 := l2
	ten := 0
	nums := make([]int, 0)
	for {
		tmp := cur1.Val + cur2.Val + ten
		nums = append(nums, tmp%10)
		ten = tmp / 10

		if cur1.Next == nil && cur2.Next == nil && ten == 0 {
			break
		}
		if cur1.Next == nil {
			cur1.Val = 0
		} else {
			cur1 = cur1.Next
		}
		if cur2.Next == nil {
			cur2.Val = 0
		} else {
			cur2 = cur2.Next
		}
	}

	var ans *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		ans = &ListNode{
			Val:  nums[i],
			Next: ans,
		}
	}

	return ans
}
