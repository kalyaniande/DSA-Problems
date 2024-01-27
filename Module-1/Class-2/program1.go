// Given an array A of N integers.
// Count the number of elements that have at least 1 elements greater than itself.

// Problem Constraints
// 1 <= N <= 10 pow 5
// 1 <= A[i] <= 10 pow 9

// Example Input
// Input 1:
// A = [3, 1, 2]
// Input 2:
// A = [5, 5, 3]

// Example Output
// Output 1:
// 2
// Output 2:
// 1

// Example Explanation
// Explanation 1:
// The elements that have at least 1 element greater than itself are 1 and 2
// Explanation 2:
// The elements that have at least 1 element greater than itself is 3

package main

import (
	"fmt"
)

func countElements(n []int) int {
	count := 0
	max := n[0]
	for i := 0; i < len(n); i++ {
		if max < n[i] {
			max = n[i]
			count++
		} else if max != n[i] {
			count++
		}
	}
	return count
}

func main() {
	n := []int{1, 1, 1}
	fmt.Println("Count:", countElements(n))
}
