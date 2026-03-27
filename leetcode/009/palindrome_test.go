package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert.True(t, isPalindrome(121), "Это палиндром")
	assert.False(t, isPalindrome(-121), "Отрицательное число не палиндром")
	assert.False(t, isPalindrome(10), "Лидирующий ноль не палиндром")
}
