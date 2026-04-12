package main

func main() {
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
