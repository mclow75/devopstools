package main

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

}
