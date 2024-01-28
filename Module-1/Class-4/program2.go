// Given an integer array A containing N distinct integers, you have to find all the leaders in array A.
// An element is a leader if it is strictly greater than all the elements to its right side.

// NOTE: The rightmost element is always a leader.

// Problem Constraints
// 1 <= N <= 10 pow 5
// 1 <= A[i] <= 10 pow 8

// Example Input
// Input 1:
// A = [16, 17, 4, 3, 5, 2]

// Input 2:
//  A = [5, 4]

// Example Output
// Output 1:
// [17, 2, 5]

// Output 2:
// [5, 4]

// Example Explanation
// Explanation 1:
// Element 17 is strictly greater than all the elements on the right side to it.
//  Element 2 is strictly greater than all the elements on the right side to it.
//  Element 5 is strictly greater than all the elements on the right side to it.
//  So we will return these three elements i.e [17, 2, 5], we can also return [2, 5, 17] or [5, 2, 17] or any other ordering.

// Explanation 2:
// Element 5 is strictly greater than all the elements on the right side to it.
//  Element 4 is strictly greater than all the elements on the right side to it.
//  So we will return these two elements i.e [5, 4], we can also any other ordering.

package main

import (
	"fmt"
)

func solve(A []int) []int {
	output := []int{}
	max := A[len(A)-1]
	for i := len(A) - 1; i >= 0; i-- {
		if i == len(A)-1 {
			output = append(output, A[i])
		} else if A[i] > max {
			output = append(output, A[i])
			max = A[i]
		}
	}
	return output
}

func main() {
	A := []int{16, 17, 4, 3, 5, 2}
	fmt.Println("Output: ", solve(A))
}
