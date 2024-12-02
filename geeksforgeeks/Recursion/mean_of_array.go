package main

import (
	"fmt"
)

func mean(A []int, N int) int {
	if N <= 1 {
		return 0
	}
	if N == 1 {
		return A[0]
	}

	return ((mean(A, N-1)*(N-1) + A[N-1]) / N)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	size := len(arr)

	m := mean(arr, size)

	fmt.Println("mean: ", m)
}
