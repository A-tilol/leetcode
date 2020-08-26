package main

var g [][]byte
var visited [][]bool
var m, n int

func dfs(i, j int) {
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	if g[i][j] != '1' || visited[i][j] {
		return
	}
	visited[i][j] = true

	dfs(i-1, j)
	dfs(i+1, j)
	dfs(i, j-1)
	dfs(i, j+1)
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	g = grid
	m, n = len(grid), len(grid[0])

	visited = make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '1' || visited[i][j] {
				continue
			}
			dfs(i, j)
			ans++
		}
	}

	return ans
}

func main() {
}
