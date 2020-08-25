package main

import "fmt"

var b [][]byte
var m, n int
var ws string
var dirs = [][]int{
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
	[]int{-1, 0},
}

func dfs(i, j, k int) bool {
	if k == len(ws) {
		return true
	}
	if i < 0 || i >= m || j < 0 || j >= n || b[i][j] != ws[k] {
		return false
	}
	b[i][j] = 0

	for _, dir := range dirs {
		if dfs(i+dir[0], j+dir[1], k+1) {
			return true
		}
	}
	b[i][j] = ws[k]

	return false
}

// no aditional space
func exist(board [][]byte, word string) bool {
	b, ws = board, word
	m, n = len(board), len(board[0])
	for i := range board {
		for j := range board[0] {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs1(i, j, k int, memo []bool) bool {
	if memo[i*n+j] {
		return false
	}
	if k == len(ws)-1 {
		return true
	}
	memo[i*n+j] = true

	for _, dir := range dirs {
		ni, nj := i+dir[0], j+dir[1]
		if ni >= 0 && ni < m && nj >= 0 && nj < n && b[ni][nj] == ws[k+1] {
			memocopy := make([]bool, m*n)
			copy(memocopy, memo)
			done := dfs1(ni, nj, k+1, memocopy)
			if done {
				return done
			}
		}
	}

	return false
}

func exist1(board [][]byte, word string) bool {
	b = board
	ws = word
	m, n = len(board), len(board[0])

	for i := range board {
		for j := range board[0] {
			if board[i][j] == word[0] {
				found := dfs1(i, j, 0, make([]bool, m*n))
				if found {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	fmt.Println(exist([][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}, "ABCCED"))
	fmt.Println(exist([][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}, "SEE"))
	fmt.Println(exist([][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}, "ABCB"))
}
