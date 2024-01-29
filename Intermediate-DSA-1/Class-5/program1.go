// Given an array A of length N, return the subarray from B to C.

// Problem Constraints
// 1 <= N <= 10 pow 5
// 1 <= A[i] <= 10 pow 9
// 0 <= B <= C < N

// Example Input
// Input 1:
// A = [4, 3, 2, 6]
// B = 1
// C = 3

// Input 2:
//  A = [4, 2, 2]
// B = 0
// C = 1

// Example Output
// Output 1:
// [3, 2, 6]

// Output 2:
// [4, 2]

// Example Explanation
// Explanation 1:
// The subarray of A from 1 to 3 is [3, 2, 6].

// Explanation 2:
// The subarray of A from 0 to 1 is [4, 2].

package main

import (
	"fmt"
)

func solve(A []int, B int, C int) []int {
	output := make([]int, (C - B + 1))
	j := 0
	for i := B; i <= C; i++ {
		output[j] = A[i]
		j++
	}
	return output
}

func main() {
	A := []int{4, 3, 2, 6}
	B := 1
	C := 3
	fmt.Println("Output: ", solve(A, B, C))
}
