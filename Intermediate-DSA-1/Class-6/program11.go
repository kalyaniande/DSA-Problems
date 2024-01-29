// You are given two integer matrices A and B having same size(Both having same number of rows (N) and columns (M).
// You have to subtract matrix B from A and return the resultant matrix. (i.e. return the matrix A - B).
// If A and B are two matrices of the same order (same dimensions).
// Then A - B is a matrix of the same order as A and B and its elements are obtained by doing an element wise subtraction of A from B.

// Problem Constraints
// 1 <= N, M <= 10 pow 3
// -10 pow 9 <= A[i][j], B[i][j] <= 10 pow 9

// Example Input
// Input 1:
// A =  [[1, 2, 3],
//       [4, 5, 6],
//       [7, 8, 9]]

// B =  [[9, 8, 7],
//       [6, 5, 4],
//       [3, 2, 1]]

// Input 2:
// A = [[1, 1]]

// B = [[2, 3]]

// Example Output
// Output 1:
// [[-8, -6, -4],
//   [-2, 0, 2],
//   [4, 6, 8]]

// Output 2:
//  [[-1, -2]]

package main

import (
	"fmt"
)

func solve(A [][]int, B [][]int) [][]int {
	output := A
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			output[i][j] = A[i][j] - B[i][j]
		}
	}
	return output
}

func main() {
	A := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	B := [][]int{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}}
	fmt.Println("Output: ", solve(A, B))
}

// Time Complexity: O(N*M)
// Space Complexity: O(1)
