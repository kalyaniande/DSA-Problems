// You are given a N X N integer matrix. You have to find the sum of all the minor diagonal elements of A.
// Minor diagonal of a M X M matrix A is a collection of elements A[i, j] such that i + j = M + 1 (where i, j are 1-based).

// Problem Constraints
// 1 <= N <= 10 pow 3
// -1000 <= A[i][j] <= 1000

// Example Input
// Input 1:
// A = [[1, -2, -3],
// [-4, 5, -6],
// [-7, -8, 9]]

// Input 2:
// A = [[3, 2],
// [2, 3]]

// Example Output
// Output 1:
// -5

// Output 2:
// 4

// Example Explanation
// Explanation 1:
// A[1][3] + A[2][2] + A[3][1] = (-3) + 5 + (-7) = -5

// Explanation 2:
//  A[1][2] + A[2][1] = 2 + 2 = 4

package main

import (
	"fmt"
)

func solve(A [][]int) int {
	n := len(A)
	output := 0
	j := len(A) - 1
	for i := 0; i < n; i++ {
		output += A[i][j]
		j--
	}

	return output
}

func main() {
	A := [][]int{{1, -2, -3}, {-4, 5, -6}, {-7, -8, 9}}
	fmt.Println("Output: ", solve(A))
}
