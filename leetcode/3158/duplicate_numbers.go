package main

import "fmt"

func main() {
	var test1 = []int{1, 2, 1, 3} // 1
	var test2 = []int{1, 2, 3}    // 0
	var test3 = []int{1, 2, 2, 1} // 3
	fmt.Println(singleNumber(test1))
	fmt.Println(singleNumber(test2))
	fmt.Println(singleNumber(test3))
}

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = (ones ^ num) &^ twos
		twos = (twos ^ num) &^ ones
		fmt.Println("data:", ones, twos)
	}

	//	if ones^twos == 3 {
	//		return ones ^ twos
	//	}
	return ones
}
