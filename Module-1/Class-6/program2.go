// You are given a 2D integer matrix A, return a 1D integer array containing row-wise sums of original matrix.

// Problem Constraints
// 1 <= A.size() <= 10 pow 3
// 1 <= A[i].size() <= 10 pow 3
// 1 <= A[i][j] <= 10 pow 3

// Example Input
// Input 1:
// [1,2,3,4]
// [5,6,7,8]
// [9,2,3,4]

// Example Output
// Output 1:
// [10,26,18]

// Example Explanation
// Explanation 1:
// Row 1 = 1+2+3+4 = 10
// Row 2 = 5+6+7+8 = 26
// Row 3 = 9+2+3+4 = 18

package main

import (
	"fmt"
)

func solve(A [][]int) []int {
	row_count := len(A)
	column_count := len(A[0])
	output := make([]int, row_count)

	for i := 0; i < row_count; i++ {
		sum := 0
		for j := 0; j < column_count; j++ {
			sum += A[i][j]
		}
		output[i] = sum
	}

	return output
}

func main() {
	A := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 2, 3, 4}}
	fmt.Println("Output: ", solve(A))
}
