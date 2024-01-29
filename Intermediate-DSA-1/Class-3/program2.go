// Given an array A of N integers. Construct prefix sum of the array in the given array itself.

// Problem Constraints
// 1 <= N <= 10 pow 5
// 1 <= A[i] <= 10 pow 3

// Example Input
// Input 1:
// A = [1, 2, 3, 4, 5]

// Input 2:
// A = [4, 3, 2]

// Example Output
// Output 1:
// [1, 3, 6, 10, 15]

// Output 2:
// [4, 7, 9]

// Example Explanation
// Explanation 1:
// The prefix sum array of [1, 2, 3, 4, 5] is [1, 3, 6, 10, 15].

// Explanation 2:
// The prefix sum array of [4, 3, 2] is [4, 7, 9].

package main

import (
	"fmt"
)

func prefixSum(A []int) []int {
	for i := 1; i < len(A); i++ {
		A[i] = A[i-1] + A[i]
	}
	return A
}

func main() {
	A := []int{1, 2, 3, 4, 5}
	fmt.Println("prefixsum", prefixSum(A))
}
