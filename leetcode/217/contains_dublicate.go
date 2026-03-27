package main

import "fmt"

func containsDuplicate(nums []int) bool {
	type void struct{}

	var member void

	intSet := make(map[int]void, len(nums))

	for _, num := range nums {
		_, exists := intSet[num]
		if exists {
			return true
		}
		intSet[num] = member
	}
	return false
}

func main() {
	var (
		test1 = []int{1, 2, 3, 1}
		test2 = []int{1, 2, 3, 4}
		test3 = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	)
	fmt.Println(containsDuplicate(test1))
	fmt.Println(containsDuplicate(test2))
	fmt.Println(containsDuplicate(test3))
}
