// You are given a n x n 2D matrix A representing an image.
// Rotate the image by 90 degrees (clockwise).
// You need to do this in place.

// Note: If you end up using an additional array, you will only receive partial score.

// Problem Constraints
// 1 <= n <= 1000

// Example Input
// Input 1:
// [
//     [1, 2],
//     [3, 4]
//  ]

// Input 2:
//  [
//     [1, 2, 3],
//     [4, 5, 6],
//     [7, 8, 9]
//  ]

// Example Output
// Output 1:
// [
//     [3, 1],
//     [4, 2]
//  ]

// Output 2:
//  [
//     [7, 4, 1],
//     [8, 5, 2],
//     [9, 6, 3]
//  ]

// Example Explanation
// Explanation 1:
// After rotating the matrix by 90 degree:
//  1 goes to 2, 2 goes to 4
//  4 goes to 3, 3 goes to 1

// Explanation 2:
// After rotating the matrix by 90 degree:
//  1 goes to 3, 3 goes to 9
//  2 goes to 6, 6 goes to 8
//  9 goes to 7, 7 goes to 1
//  8 goes to 4, 4 goes to 2

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

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A)/2; j++ {
			A[i][j], A[i][len(A)-1-j] = A[i][len(A)-1-j], A[i][j]
		}
	}
	return A
}

func main() {
	A := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	fmt.Println("Output: ", solve(A))
}

// Time Complexity: O(N*N)
// Space Complexity: O(1)
