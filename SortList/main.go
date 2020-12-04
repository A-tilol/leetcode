package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

// ListNode .
type ListNode struct {
	Val  int
	Next *ListNode
}

func dfs(node *ListNode, length int) *ListNode {
	// fmt.Println("val", node.Val, "len", length)
	if length == 1 {
		node.Next = nil
		return node
	}

	if length == 2 {
		if node.Val <= node.Next.Val {
			node.Next.Next = nil
			return node
		}
		newHead := node.Next
		node.Next = nil
		newHead.Next = node
		return newHead
	}

	secondNode := node
	for i := 0; i < length/2; i++ {
		secondNode = secondNode.Next
	}

	sorted2 := dfs(secondNode, length-length/2)
	sorted1 := dfs(node, length/2)

	tmpHead := &ListNode{}
	cur := tmpHead
	for sorted1 != nil || sorted2 != nil {
		val1, val2 := math.MaxInt32, math.MaxInt32
		if sorted1 != nil {
			val1 = sorted1.Val
		}
		if sorted2 != nil {
			val2 = sorted2.Val
		}
		if val1 <= val2 {
			cur.Next = sorted1
			cur = cur.Next
			sorted1 = sorted1.Next
		} else {
			cur.Next = sorted2
			cur = cur.Next
			sorted2 = sorted2.Next
		}
	}

	return tmpHead.Next
}

func sortList(head *ListNode) *ListNode {
	l := 0
	cur := head
	for cur != nil {
		l++
		cur = cur.Next
	}
	// log.Println("len", l)

	if l < 2 {
		return head
	}

	return dfs(head, l)
}

func main() {
}

var (
	readString func() string
	readBytes  func() []byte
)

func init() {
	log.SetFlags(log.Lshortfile)
	readString, readBytes = newReadString(os.Stdin)
}

func newReadString(ior io.Reader) (func() string, func() []byte) {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+11))
	r.Split(bufio.ScanWords)

	f1 := func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
	f2 := func() []byte {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Bytes()
	}
	return f1, f2
}

func readInt64() int64 {
	i, err := strconv.ParseInt(readString(), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func readInt() int {
	return int(readInt64())
}

func readFloat64() float64 {
	f, err := strconv.ParseFloat(readString(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sum(a []int) int {
	var ret int
	for i := range a {
		ret += a[i]
	}
	return ret
}

func sumFloat64(a []float64) float64 {
	var ret float64
	for i := range a {
		ret += a[i]
	}
	return ret
}

func gcd(m, n int) int {
	for m%n != 0 {
		m, n = n, m%n
	}
	return n
}

func lcm(m, n int) int {
	return m / gcd(m, n) * n
}
