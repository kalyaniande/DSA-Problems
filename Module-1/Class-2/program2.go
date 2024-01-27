// Given an array A and an integer B.
// A pair(i, j) in the array is a good pair if i != j and (A[i] + A[j] == B).
// Check if any good pair exist or not.
// Return 1 if good pair exist otherwise return 0.

// Problem Constraints
// 1 <= A.size() <= 10 pow 4
// 1 <= A[i] <= 10 pow 9
// 1 <= B <= 10 pow 9

// Example Input
// Input 1:
// A = [1,2,3,4]
// B = 7

// Input 2:
// A = [1,2,4]
// B = 4

// Input 3:
// A = [1,2,2]
// B = 4

// Example Output
// Output 1:
// 1

// Output 2:
// 0

// Output 3:
// 1

// Example Explanation
// Explanation 1:
//  (i,j) = (3,4)
// Explanation 2:

// No pair has sum equal to 4.
// Explanation 3:
//  (i,j) = (2,3)

package main

import (
	"fmt"
)

func goodPairExists(A []int, B int) int {
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if i != j {
				if A[i]+A[j] == B {
					return 1
				}
			}
		}
	}
	return 0
}

func main() {
	A := []int{1, 2, 3, 4}
	B := 5
	output := goodPairExists(A, B)
	if output == 1 {
		fmt.Println("Good Pair exists.")
	} else {
		fmt.Println("Good Pair not exists.")
	}
}
