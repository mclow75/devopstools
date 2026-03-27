package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("1.USA\n", "2.leetcode\n", "3.Google\n")
	fmt.Print(detectCapitalUse("USA"), detectCapitalUse("leetcode"), detectCapitalUse("Google"))
}

func detectCapitalUse(word string) bool {
	return (word == strings.ToLower(word) || word == strings.ToUpper(word) || (word[0:1] == strings.ToUpper(word[0:1]) && word[1:] == strings.ToLower(word[1:])))
}
