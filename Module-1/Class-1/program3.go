// Given a number A. Return square root of the number if it is perfect square otherwise return -1.

// Note: A number is a perfect square if its square root is an integer.

// Problem Constraints
// 1 <= A <= 10 pow 8

// Example Input
// Input 1:
// A = 4
// Input 2:
// A = 1001

// Example Output
// Output 1:
// 2
// Output 2:
// -1
package main

import (
	"fmt"
)

func sqrtOfNumber(n int) int {
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return i
		}
	}

	return -1
}

func main() {
	n := 101
	fmt.Println(" Square root of a number ", n, ":", sqrtOfNumber(n))
}

// Time Complexity: O(N)
// Space Complexity: O(1)
