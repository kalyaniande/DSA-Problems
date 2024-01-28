// Given an array A of N non-negative numbers and a non-negative number B,
// you need to find the number of subarrays in A with a sum less than B.
// We may assume that there is no overflow.

// Problem Constraints
// 1 <= N <= 5 x 10 pow 3
// 1 <= A[i] <= 1000
// 1 <= B <= 10 pow 7

// Example Input
// Input 1:
// A = [2, 5, 6]
//  B = 10

// Input 2:
// A = [1, 11, 2, 3, 15]
//  B = 10

// Example Output
// Output 1:
// 4

// Output 2:
// 4

// Example Explanation
// Explanation 1:
// The subarrays with sum less than B are {2}, {5}, {6} and {2, 5},

// Explanation 2:
// The subarrays with sum less than B are {1}, {2}, {3} and {2, 3}

package main

import (
	"fmt"
)

func solve(A []int, B int) int {
	count := 0
	for i := 0; i < len(A); i++ {
		sum := 0
		for j := i; j < len(A); j++ {
			sum += A[j]
			if sum < B {
				count++
			}
		}
	}

	return count
}

func main() {
	A := []int{1, 11, 2, 3, 16}
	B := 10
	fmt.Println("Output: ", solve(A, B))
}
