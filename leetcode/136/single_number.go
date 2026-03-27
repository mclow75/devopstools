package main

import "fmt"

func main() {
	var test1 = []int{2, 2, 1}
	var test2 = []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(test1))
	fmt.Println(singleNumber(test2))
}

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}
