package main

import "fmt"

func main() {
	var test1 = []int{1, 8, 6, 2, 5, 4, 8, 3, 7} // 49
	var test2 = []int{1, 1}                      // 1
	fmt.Println(maxArea(test1))
	fmt.Println(maxArea(test2))
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	area := 0

	for left < right {
		h := min(height[left], height[right])
		area = max(area, (right-left)*h)
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return area
}
