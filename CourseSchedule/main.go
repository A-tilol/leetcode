package main

import "fmt"

var visited []int
var cToPres map[int][]int

func dfs(i int) bool {
	if visited[i] == 1 {
		return true
	}
	if visited[i] == -1 {
		return false
	}
	visited[i] = -1

	depends := cToPres[i]
	for j := range depends {
		if !dfs(depends[j]) {
			return false
		}
	}
	visited[i] = 1

	return true
}

// algorithm: topological sort
func canFinish(numCourses int, prerequisites [][]int) bool {
	cToPres = make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		cToPres[i] = make([]int, 0)
	}
	for i := range prerequisites {
		p0, p1 := prerequisites[i][0], prerequisites[i][1]
		cToPres[p0] = append(cToPres[p0], p1)
	}

	visited = make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if visited[i] == 1 {
			continue
		}
		if !dfs(i) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canFinish(2, [][]int{[]int{1, 0}, []int{0, 1}}))
}
