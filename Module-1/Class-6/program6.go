// Given a 2D integer array A, return the transpose of A.
// The transpose of a matrix is the matrix flipped over its main diagonal,
// switching the matrix's row and column indices.

// Problem Constraints
// 1 <= A.size() <= 1000
// 1 <= A[i].size() <= 1000
// 1 <= A[i][j] <= 1000

// Example Input
// Input 1:
// A = [[1, 2, 3],[4, 5, 6],[7, 8, 9]]

// Input 2:
// A = [[1, 2],[1, 2],[1, 2]]

// Example Output
// Output 1:
// [[1, 4, 7], [2, 5, 8], [3, 6, 9]]

// Output 2:
// [[1, 1, 1], [2, 2, 2]]

// Example Explanation
// Explanation 1:
// Clearly after converting rows to column and columns to rows of [[1, 2, 3],[4, 5, 6],[7, 8, 9]]
//  we will get [[1, 4, 7], [2, 5, 8], [3, 6, 9]].

// Explanation 2:
// After transposing the matrix, A becomes [[1, 1, 1], [2, 2, 2]]

package main

import (
	"fmt"
)

func solve(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		j := i
		for j < len(A) {
			A[i][j], A[j][i] = A[j][i], A[i][j]
			j++
		}

	}
	return A
}

func main() {
	A := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Output: ", solve(A))
}

// Time Complexity: O(N*M)
// Space Complexity: O(1)
