// Given an array A and an integer B, find the number of occurrences of B in A.

// Problem Constraints
// 1 <= B, Ai <= 10 pow 9
// 1 <= length(A) <= 10 pow 5

// Example Input
// Input 1:
// A = [1, 2, 2], B = 2

// Input 2:
// A = [1, 2, 1], B = 3

// Example Output
// Output 1:
// 2

// Output 2:
// 0

package main

import (
	"fmt"
)

func numOfOccurances(A []int, B int) int {
	count := 0
	for i := 0; i < len(A); i++ {
		if B == A[i] {
			count++
		}
	}
	return count
}

func main() {
	A := []int{1, 2, 2}
	B := 2
	fmt.Println("Number of Occurances:", numOfOccurances(A, B))
}
