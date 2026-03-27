package main

import "fmt"

func main() {
	ints := []int{1, 2, 3}
	var intSum int = sum(ints)

	fmt.Println("Integer: ", intSum)

	fl := []float32{2.3, 3, 4.7, 9.1}
	var floatSum float32 = sum(fl)
	fmt.Println("Float: ", floatSum)
}

func sum[T int | float32](input []T) T {
	var s T
	for _, i := range input {
		s += i
	}
	return s
}
