// You are given two matrices A & B of same size, you have to return another matrix which is the sum of A and B.
// Note: Matrices are of same size means the number of rows and number of columns of both matrices are equal.

// Problem Constraints
// 1 <= A.size(), B.size() <= 1000 1 <= A[i].size(), B[i].size() <= 1000 1 <= A[i][j], B[i][j] <= 1000

// Example Input
// Input 1:
// A = [[1, 2, 3],
// [4, 5, 6],
// [7, 8, 9]]

// B = [[9, 8, 7],
// [6, 5, 4],
// [3, 2, 1]]

// Input 2:
// A = [[1, 2, 3],
// [4, 1, 2],
// [7, 8, 9]]

// B = [[9, 9, 7],
// [1, 2, 4],
// [4, 6, 3]]

// Example Output
// Output 1:
// [[10, 10, 10],
//  [10, 10, 10],
//  [10, 10, 10]]

// Output 2:
//  [[10, 11, 10],
//  [5,   3,  6],
//  [11, 14, 12]]

// Example Explanation
// Explanation 1:
// A + B = [[1+9, 2+8, 3+7],
// [4+6, 5+5, 6+4],
// [7+3, 8+2, 9+1]]
// = [[10, 10, 10],
// [10, 10, 10],
// [10, 10, 10]].

// Explanation 2:
// A + B = [[1+9, 2+9, 3+7],
// [4+1, 1+2, 2+4],
// [7+4, 8+6, 9+3]]
// = [[10, 11, 10],
// [5,   3,  6],
// [11, 14, 12]].

package main

import (
	"fmt"
)

func solve(A [][]int, B [][]int) [][]int {
	output := A
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			output[i][j] = A[i][j] + B[i][j]
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
