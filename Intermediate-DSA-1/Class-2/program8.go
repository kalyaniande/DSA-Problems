// You are given an integer array A.
// You have to find the second largest element/value in the array or report that no such element exists.
// Return the second largest element. If no such element exist then return -1.

// Problem Constraints
// 1 <= |A| <= 10 pow 5
// 0 <= A[i] <= 10 pow 9

// Example Input
// Input 1:
//  A = [2, 1, 2]

// Input 2:
// A = [2]

// Example Output
// Output 1:
// 1

// Output 2:
// -1

package main

import (
	"fmt"
)

func secondLargest(A []int) int {
	if len(A) < 2 {
		return -1
	}
	max := A[0]
	output := A[0]
	for i := 0; i < len(A); i++ {
		if max < A[i] {
			max = A[i]
		}
	}

	for i := 0; i < len(A); i++ {
		if output < A[i] && output < max {
			output = A[i]
		} else if output == max {
			output = A[i]
		}
	}
	return output

}

func main() {
	A := []int{1, 8, 3}
	fmt.Println("Second Largest Element:", secondLargest(A))
}
