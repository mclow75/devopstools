package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// Create HaspMap
	hm := make(map[int]int, len(nums))
	// Save to hashmap
	for i, num := range nums {
		_, ok := hm[num]
		if ok {
			return []int{hm[num], i}
		}
		hm[target-num] = i
	}
	return nil
}

func main() {
	var test1 = []int{2, 7, 11, 15}
	fmt.Println(twoSum(test1, 9))
}
