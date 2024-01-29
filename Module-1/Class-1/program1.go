// Given an integer A, you need to find the count of it's factors.

// Factor of a number is the number which divides it perfectly leaving no remainder.

// Example : 1, 2, 3, 6 are factors of 6

// Problem Constraints
// 1 <= A <= 10 pow 9

// Example Input
// Input 1:
// 5
// Input 2:
// 10

// Example Output
// Output 1:
// 2
// Output 2:
// 4
package main

import (
	"fmt"
)

func countFactors(n int) int {
	count := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i == n/i {
				count++
			} else {
				count = count + 2
			}
		}
	}

	return count
}

func main() {
	n := 100
	fmt.Println(n, " factors count:", countFactors(n))
}

// Time Complexity: O(N)
// Space Complexity: O(1)
