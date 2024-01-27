// Given an array A of size N.
// You need to find the sum of Maximum and Minimum element in the given array.

// Problem Constraints
// 1 <= N <= 10 pow 5
// -10 pow 9 <= A[i] <= 10 pow 9

// Example Input
// Input 1:
// A = [-2, 1, -4, 5, 3]

// Input 2:
// A = [1, 3, 4, 1]

// Example Output
// Output 1:
// 1

// Output 2:
// 5

package main

import (
	"fmt"
)

func minMaxSum(A []int) int {
	min := A[0]
	max := A[0]
	for i := 0; i < len(A); i++ {
		if min > A[i] {
			min = A[i]
		}
		if max < A[i] {
			max = A[i]
		}
	}
	return min + max
}

func main() {
	A := []int{-2, 1, -4, 5, 3}
	fmt.Println("Min Max Sum:", minMaxSum(A))
}
