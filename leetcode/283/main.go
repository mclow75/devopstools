package main

import "fmt"

func main() {
	var (
		test1 = []int{1, 2, 3, 1}
		test2 = []int{0}
		test3 = []int{0, 1, 0, 3, 12}
	)
	fmt.Println(moveZeroes(test1))
	fmt.Println(moveZeroes(test2))
	fmt.Println(moveZeroes(test3))
}

func moveZeroes(nums []int) []int {
	l := 0

	for r := range len(nums) {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l += 1
		}
	}

	return nums
}
