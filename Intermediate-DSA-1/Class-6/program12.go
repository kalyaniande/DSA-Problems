// You are given a 2D integer matrix A, make all the elements in a row or column zero if the A[i][j] = 0.
// Specifically, make entire ith row and jth column zero.

// Problem Constraints
// 1 <= A.size() <= 10 pow 3
// 1 <= A[i].size() <= 10 pow 3
// 0 <= A[i][j] <= 10 pow 3

// Example Input
// Input 1:
// [1,2,3,4]
// [5,6,7,0]
// [9,2,0,4]

// Example Output
// Output 1:
// [1,2,0,0]
// [0,0,0,0]
// [0,0,0,0]

// Example Explanation
// Explanation 1:
// A[2][4] = A[3][3] = 0, so make 2nd row, 3rd row, 3rd column and 4th column zero.

package main

import (
	"fmt"
)

func solve(A [][]int) [][]int {
	indexes := [][]int{}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] == 0 {
				indexes = append(indexes, []int{i, j})
			}
		}
	}

	for i := 0; i < len(indexes); i++ {
		for j := 0; j < len(A[indexes[i][0]]); j++ {
			A[indexes[i][0]][j] = 0
		}

		for k := 0; k < len(A); k++ {
			A[k][indexes[i][1]] = 0
		}
	}
	fmt.Println("indexes", indexes)
	return A
}

func main() {
	A := [][]int{{1, 2, 3, 4}, {5, 6, 7, 0}, {9, 2, 0, 4}}
	fmt.Println("Output: ", solve(A))
}

// Time Complexity: O(N*M*N*M)
// Space Complexity: O(N*M)
