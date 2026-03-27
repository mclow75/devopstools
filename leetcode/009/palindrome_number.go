package main

import "fmt"

func isPalindrome(x int) bool {
	orig := x
	rev := 0
	last := 0
	for orig > 0 {
		last = orig % 10
		orig = orig / 10
		rev = rev*10 + last
	}
	return rev == x
}

func main() {
	fmt.Println("121", isPalindrome(121))
	fmt.Println("-121", isPalindrome(-121))
	fmt.Println("10", isPalindrome(10))
	fmt.Println("456", isPalindrome(456))
}
