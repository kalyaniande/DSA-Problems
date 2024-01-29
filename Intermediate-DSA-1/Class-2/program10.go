// Given an integer array A of size N. In one second, you can increase the value of one element by 1.
// Find the minimum time in seconds to make all elements of the array equal.

// Problem Constraints
// 1 <= N <= 1000000
// 1 <= A[i] <= 1000

// Example Input
// A = [2, 4, 1, 3, 2]

// Example Output
// 8

// Example Explanation
// We can change the array A = [4, 4, 4, 4, 4]. The time required will be 8 seconds.

package main

import (
	"fmt"
)

func solve(A []int) int {
	output := 0
	max := A[0]
	for i := 0; i < len(A); i++ {
		if max < A[i] {
			max = A[i]
		}
	}

	for i := 0; i < len(A); i++ {
		if A[i] < max {
			output = output + (max - A[i])
		}
	}
	return output

}

func main() {
	A := []int{2, 4, 1, 3, 2}
	fmt.Println("Output:", solve(A))
}
