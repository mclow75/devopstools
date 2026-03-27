package main

import "fmt"

func main() {
	var test1 = []int{2, 2, 3, 2}
	var test2 = []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(test1))
	fmt.Println(singleNumber(test2))
}

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = (ones ^ num) &^ twos
		twos = (twos ^ num) &^ ones
	}
	return ones
}
