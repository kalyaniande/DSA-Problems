// Give a N * N square matrix A, return an array of its anti-diagonals. Look at the example for more details.

// Problem Constraints
// 1 <= N <= 10 pow 3
// 1<= A[i][j] <= 1e9

// Example Input
// Input 1:
// A = [[1, 2, 3],
// [4, 5, 6],
// [7, 8, 9]]

// Input 2:
// A = [[1, 2],
// [3, 4]]

// Example Output
// Output 1:
// 1 0 0
// 2 4 0
// 3 5 7
// 6 8 0
// 9 0 0

// Output 2:
// 1 0
// 2 3
// 4 0

// Example Explanation
// Explanation 1:
// The first anti diagonal of the matrix is [1 ], the rest spaces shoud be filled with 0 making the row as [1, 0, 0].
// The second anti diagonal of the matrix is [2, 4 ], the rest spaces shoud be filled with 0 making the row as [2, 4, 0].
// The third anti diagonal of the matrix is [3, 5, 7 ], the rest spaces shoud be filled with 0 making the row as [3, 5, 7].
// The fourth anti diagonal of the matrix is [6, 8 ], the rest spaces shoud be filled with 0 making the row as [6, 8, 0].
// The fifth anti diagonal of the matrix is [9 ], the rest spaces shoud be filled with 0 making the row as [9, 0, 0].

// Explanation 2:
//  The first anti diagonal of the matrix is [1 ], the rest spaces shoud be filled with 0 making the row as [1, 0, 0].
// The second anti diagonal of the matrix is [2, 4 ], the rest spaces shoud be filled with 0 making the row as [2, 4, 0].
// The third anti diagonal of the matrix is [3, 0, 0 ], the rest spaces shoud be filled with 0 making the row as [3, 0, 0].

package main

import (
	"fmt"
)

func solve(A [][]int) [][]int {
	var output [][]int
	length := len(A)

	for j := 0; j < len(A[0]); j++ {
		x := 0
		y := j
		tmp := make([]int, length)
		for x < len(A) && y >= 0 {
			tmp[x] = A[x][y]
			x++
			y--
		}
		output = append(output, tmp)
	}

	for i := 1; i < len(A); i++ {
		x := i
		y := len(A) - 1
		index := 0
		tmp := make([]int, length)
		for x < len(A) && y >= 0 {
			tmp[index] = A[x][y]
			x++
			y--
			index++
		}
		output = append(output, tmp)
	}

	return output
}

func main() {
	A := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Output: ", solve(A))
}
