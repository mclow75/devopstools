package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOne(t *testing.T) {
	actual := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	expected := 49
	assert.Equal(t, expected, actual)
}
func TestTwo(t *testing.T) {
	actual := maxArea([]int{1, 1})
	expected := 1
	assert.Equal(t, expected, actual)
}
