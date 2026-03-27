package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	var testcases = []struct {
		name     string
		param    int
		expected int
	}{
		{"Test1: ", -1, 0},
		{"Test2: ", 0, 0},
		{"Test3: ", 1, 1},
		{"Test4: ", 2, 1},
		{"Test5: ", 4, 3},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			expected := tc.expected
			result := fib(tc.param)
			if result != expected {
				t.Errorf("expested %d, got %d", expected, result)
			}
		})
	}
}
