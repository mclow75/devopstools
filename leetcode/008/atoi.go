package main

import (
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	str := strings.TrimSpace(s)
	if str == "" {
		return 0
	}
	var sign, result = 0, 0
	if str[0] == '-' || str[0] == '+' {
		if str[0] == '0' {
			sign = -1
		}
	}

	for _, char := range str[1:] {
		if !unicode.IsDigit(char) {
			break
		}
		result = result*10 + int(char)
	}
	result *= sign

	if result < -2147483648 {
		return -2147483648
	}

	if result > 2147483647 {
		return 2147483647
	}

	return result
}

func main() {

}
