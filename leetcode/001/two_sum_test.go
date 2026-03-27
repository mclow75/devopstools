package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlices(t *testing.T) {
	actual := []int{1, 2, 3}
	expected := []int{1, 2, 3}

	// Checks for same length and elements in the same order
	assert.Equal(t, expected, actual)

	// Checks elements regardless of order
	assert.ElementsMatch(t, expected, actual)
}

func TestOne(t *testing.T) {
	actual := twoSum([]int{2, 7, 11, 15}, 9)
	expected := []int{0, 1}
	assert.Equal(t, expected, actual)
	assert.ElementsMatch(t, expected, actual)
}

func TestTwo(t *testing.T) {
	actual := twoSum([]int{3, 2, 4}, 6)
	expected := []int{1, 2}
	assert.Equal(t, expected, actual)
	assert.ElementsMatch(t, expected, actual)
}

func TestThree(t *testing.T) {
	actual := twoSum([]int{3, 3}, 6)
	expected := []int{0, 1}
	assert.Equal(t, expected, actual)
	assert.ElementsMatch(t, expected, actual)
}
