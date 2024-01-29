// You are given a N X N integer matrix. You have to find the sum of all the main diagonal elements of A.

// Problem Constraints
// 1 <= N <= 10 pow 3
// -1000 <= A[i][j] <= 1000

// Example Input
// Input 1:
// [[2,7,5],[8,9,4], [3,1,6]]

// Example Output
// Output 1:
// 17

// Example Explanation
// Explanation 1:
// 2 + 9 + 6 = 17

package main

import (
	"fmt"
)

func solve(A [][]int) int {
	n := len(A)
	output := 0

	for i := 0; i < n; i++ {
		output += A[i][i]
	}

	return output
}

func main() {
	A := [][]int{{2, 7, 5}, {8, 9, 4}, {3, 1, 6}}
	fmt.Println("Output: ", solve(A))
}

// Time Complexity: O(N)
// Space Complexity: O(1)
