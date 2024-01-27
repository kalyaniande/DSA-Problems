// Given an array of integers A, Q queries given by the 2D array B of size Q×2.
// For each query calculate sum of all even indices in a given range

// Problem Constraints
// 2 <= length of the array <= 1000
// 1 <= A[i] <= 10

// Input 1:
//   A = [3, 4, -2, 8, 6, 2, 1, 3]
//   B = [[2, 6], [3, 7]]
// Output 1:
//     [5, 7]

package main

import (
	"fmt"
)

func solve(A []int, B [][]int) []int {
	output := make([]int, len(B))
	for i := 1; i < len(A); i++ {
		if i%2 == 0 {
			A[i] = A[i-1] + A[i]
		} else {
			A[i] = A[i-1]
		}
	}

	for i := 0; i < len(B); i++ {
		l := B[i][0]
		r := B[i][1]

		if l == 0 {
			output[i] = A[r]
		} else {
			output[i] = A[r] - A[l-1]
		}
	}
	return output
}

func main() {
	A := []int{3, 4, -2, 8, 6, 2, 1, 3}
	B := [][]int{{2, 6}, {3, 7}}
	fmt.Println("Output:", solve(A, B))
}
