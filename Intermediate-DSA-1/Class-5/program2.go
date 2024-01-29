// You are given an integer array C of size A.
// Now you need to find a subarray (contiguous elements) so that the sum of contiguous elements is maximum.
// But the sum must not exceed B.

// Problem Constraints
// 1 <= A <= 10 pow 3
// 1 <= B <= 10 pow 9
// 1 <= C[i] <= 10 pow 6

// Example Input
// Input 1:
// A = 5
// B = 12
// C = [2, 1, 3, 4, 5]

// Input 2:
//  A = [4, 2, 2]
// B = 0
// C = 1

// Example Output
// Output 1:
// 12

// Output 2:
// 0

// Example Explanation
// Explanation 1:
// We can select {3,4,5} which sums up to 12 which is the maximum possible sum.

// Explanation 2:
// All elements are greater than B, which means we cannot select any subarray.
// Hence, the answer is 0.

package main

import (
	"fmt"
)

func solve(C []int, B int, A int) int {
	max_sum := 0

	for i := 0; i < A; i++ {
		sum := 0
		for j := i; j < A; j++ {
			sum += C[j]
			if sum > max_sum && sum <= B {
				max_sum = sum
			}
		}
	}

	return max_sum
}

func main() {
	C := []int{2, 1, 3, 4, 5}
	B := 12
	A := 5
	fmt.Println("Output: ", solve(C, B, A))
}
