package main

// Intuition
// Since the array is sorted and all elements except one appear exactly twice,
// the duplicate elements will always be adjacent. Before the single element,
// the first occurrence of a pair will be at an even index, and the second will
//  be at an odd index. This pattern will be broken at the single element.
// A binary search can help find this point efficiently.
// Approach
// We use binary search to find the unique element. Here's the idea:
//
// Set left to 0 and right to the last index.
// Find the mid index.
// Check whether nums[mid] is equal to its neighbor (nums[mid-1] or nums[mid+1]).
// Deprighting on the parity of mid, we update left or right to eliminate half of
// the array.
// If nums[mid] is not equal to either neighbor, it must be the single element.
// This way, we achieve O(log n) time complexity instead of a linear scan.
// Complexity
// Time complexity: O(logn)
// Space complexity: O(1)

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if mid != 0 && nums[mid] == nums[mid-1] {
			if mid%2 != 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if mid != len(nums)-1 && nums[mid] == nums[mid+1] {
			if mid%2 == 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			return nums[mid]
		}
	}
	return -1
}

func main() {

}
