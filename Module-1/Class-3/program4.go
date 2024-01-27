// You are given an array A of length N and Q queries given by the 2D array B of size Q×2.
// Each query consists of two integers B[i][0] and B[i][1].
// For every query, your task is to find the count of even numbers in the range from A[B[i][0]] to A[B[i][1]].

// Problem Constraints
// 1 <= N <= 10 pow 5
// 1 <= Q <= 10 pow 5
// 1 <= A[i] <= 10 pow 9
// 0 <= B[i][0] <= B[i][1] < N

// Example Input
// Input 1:
// A = [1, 2, 3, 4, 5]
// B = [   [0, 2]
// [2, 4]
// [1, 4]   ]

// Input 2:
// A = [2, 1, 8, 3, 9, 6]
// B = [   [0, 3]
// [3, 5]
// [1, 3]
// [2, 4]   ]

// Example Output
// Output 1:
// [1, 1, 2]

// Output 2:
// [2, 1, 1, 1]

// Example Explanation
// For Input 1:
// The subarray for the first query is [1, 2, 3] (index 0 to 2) which contains 1 even number.
// The subarray for the second query is [3, 4, 5] (index 2 to 4) which contains 1 even number.
// The subarray for the third query is [2, 3, 4, 5] (index 1 to 4) which contains 2 even numbers.

// For Input 2:
// The subarray for the first query is [2, 1, 8, 3] (index 0 to 3) which contains 2 even numbers.
// The subarray for the second query is [3, 9, 6] (index 3 to 5) which contains 1 even number.
// The subarray for the third query is [1, 8, 3] (index 1 to 3) which contains 1 even number.
// The subarray for the fourth query is [8, 3, 9] (index 2 to 4) which contains 1 even number.

package main

import (
	"fmt"
)

func solve(A []int, B [][]int) []int {
	output := make([]int, len(B))
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			if i == 0 {
				A[i] = 1
			} else {
				A[i] = A[i-1] + 1
			}
		} else {
			if i == 0 {
				A[i] = 0
			} else {
				A[i] = A[i-1]
			}
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
	fmt.Println(A)
	return output
}
func main() {
	A := []int{2, 1, 8, 3, 9, 6}
	B := [][]int{{0, 3}, {3, 5}, {1, 3}, {2, 4}}
	fmt.Println("Output", solve(A, B))
}
