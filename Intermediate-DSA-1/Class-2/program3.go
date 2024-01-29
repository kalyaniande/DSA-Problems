// Given an array A of N integers and also given two integers B and C.
// Reverse the elements of the array A within the given inclusive range [B, C].
// Return the array A after reversing in the given range.

// Problem Constraints
// 1 <= N <= 10 pow 5
// 1 <= A[i] <= 10 pow 9
// 0 <= B <= C <= N - 1

// Example Input
// Input 1:
// A = [1, 2, 3, 4]
// B = 2
// C = 3

// Input 2:
// A = [2, 5, 6]
// B = 0
// C = 2

// Example Output
// Output 1:
// [1, 2, 4, 3]

// Output 2:
// [6, 5, 2]

// Example Explanation
// Explanation 1:
// We reverse the subarray [3, 4].

// Explanation 2:
// We reverse the entire array [2, 5, 6].

package main

import (
	"fmt"
)

func reverseElements(A []int, B int, C int) []int {

	for B <= C {
		A[B], A[C] = A[C], A[B]
		B++
		C--
	}
	return A
}

func main() {
	A := []int{1, 2, 3, 4}
	B := 2
	C := 3
	fmt.Println("Output:", reverseElements(A, B, C))
}
