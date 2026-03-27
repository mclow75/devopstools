package main

import "fmt"

func climbStairs(n int) int {
	twoBack, oneBack, nextNum := 1, 2, 0

	if n < 3 {
		return n
	}
	for i := 2; i < n; i++ {
		nextNum = twoBack + oneBack
		twoBack = oneBack
		oneBack = nextNum
	}
	return oneBack
}

func main() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
}
