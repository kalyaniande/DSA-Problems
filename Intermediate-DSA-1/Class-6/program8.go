// You are given a matrix A and and an integer B, you have to perform scalar multiplication of matrix A with an integer B.

// Problem Constraints
// 1 <= A.size() <= 1000
// 1 <= A[i].size() <= 1000
// 1 <= A[i][j] <= 1000
// 1 <= B <= 1000

// Example Input
// Input 1:
// A = [[1, 2, 3],
//      [4, 5, 6],
//      [7, 8, 9]]
// B = 2

// Input 2:
// A = [[1]]
// B = 5

// Example Output
// Output 1:
// [[2, 4, 6],
// [8, 10, 12],
// [14, 16, 18]]

// Output 2:
//  [[5]]

// Example Explanation
// Explanation 1:
// ==> ( [[1, 2, 3],[4, 5, 6],[7, 8, 9]] ) * 2

// ==> [[2*1, 2*2, 2*3],
// [2*4, 2*5, 2*6],
// [2*7, 2*8, 2*9]]

// ==> [[2,   4,  6],
// [8,  10, 12],
// [14, 16, 18]]

// Explanation 2:
// ==> ( [[1]] ) * 5
// ==> [[5*1]]
// ==> [[5]]

package main

import (
	"fmt"
)

func solve(A [][]int, B int) [][]int {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			A[i][j] = A[i][j] * B
		}
	}
	return A
}

func main() {
	A := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	B := 2
	fmt.Println("Output: ", solve(A, B))
}

// Time Complexity: O(N*M)
// Space Complexity: O(1)
