// You are given an integer array A of length N.
// You are also given a 2D integer array B with dimensions M x 2, where each row denotes a [L, R] query.
// For each query, you have to find the sum of all elements from L to R indices in A (0 - indexed).
// More formally, find A[L] + A[L + 1] + A[L + 2] +... + A[R - 1] + A[R] for each query.

// Problem Constraints
// 1 <= N, M <= 10 pow 5
// 1 <= A[i] <= 10 pow 9
// 0 <= L <= R < N

// Example Input
// Input 1:
// A = [1, 2, 3, 4, 5]
// B = [[0, 3], [1, 2]]

// Input 2:
// A = [2, 2, 2]
// B = [[0, 0], [1, 2]]

// Example Output
// Output 1:
// [10, 5]

// Output 2:
// [2, 4]

// Example Explanation
// Explanation 1:
// The sum of all elements of A[0 ... 3] = 1 + 2 + 3 + 4 = 10.
// The sum of all elements of A[1 ... 2] = 2 + 3 = 5.

// Explanation 2:
// The sum of all elements of A[0 ... 0] = 2 = 2.
// The sum of all elements of A[1 ... 2] = 2 + 2 = 4.

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

func solve(A []int, B [][]int) []int {
	output := make([]int, len(B))

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
	A := []int{1, 2, 3, 4, 5}
	B := [][]int{{0, 3}, {1, 2}}
	prefix_sum_arr := prefixSum(A)
	fmt.Println("prefixsum", prefix_sum_arr)
	fmt.Println("Output:", solve(prefix_sum_arr, B))
}
