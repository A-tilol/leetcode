package main

import (
	"fmt"
	"math"
)

func generateDivisors(n int) []int {
	divisors := make([]int, 0)
	sqrn := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrn; i++ {
		if n%i != 0 {
			continue
		}

		divisors = append(divisors, i)

		if i != n/i && n/i < n {
			divisors = append(divisors, n/i)
		}
	}
	return divisors
}

func divisorGame(N int) bool {
	win := make([]bool, N+1)
	win[1] = false
	for i := 2; i <= N; i++ {
		for _, d := range generateDivisors(i) {
			// fmt.Println(i, d, i-d, win[i-d])
			if !win[i-d] {
				// fmt.Println("in")
				win[i] = true
				break
			}
		}
	}

	return win[N]
}

func main() {
	fmt.Println(divisorGame(2))
	fmt.Println(divisorGame(3))
	fmt.Println(divisorGame(10))
	fmt.Println(divisorGame(1000))
}
