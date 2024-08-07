package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// T là tham số kiểu dữ liệu
// any là kiểu dữ liệu không xác định
// s là slice chứa các phần tử kiểu dữ liệu T
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func Min[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		panic("Empty slice")
	}

	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	// use PrintSlice
	// intSlice := []int{1,2,3}
	// stringSlice := []string{"thong","thuong","thai"}
	// PrintSlice[int](intSlice)
	// PrintSlice[string](stringSlice)

	// use Min
	intSlice := []int{1,2,3}
	floatSlice := []float64{1.1,2.2,3.3}
	fmt.Println(Min[int](intSlice))
	fmt.Println(Min[float64](floatSlice))
}