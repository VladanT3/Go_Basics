package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice(intSlice))

	var float32Slice = []float32{4, 5, 6}
	fmt.Println(sumSlice(float32Slice))

	var float64Slice = []float64{7, 8, 9}
	fmt.Println(sumSlice(float64Slice))

	var emptySlice = []int{}
	fmt.Println(isEmpty(emptySlice))
	fmt.Println(isEmpty(float32Slice))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for i := range slice {
		sum += slice[i]
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
