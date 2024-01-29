// You are given an integer A. You have to tell whether it is a perfect number or not.

// Perfect number is a positive integer which is equal to the sum of its proper positive divisors.

// A proper divisor of a natural number is the divisor that is strictly less than the number.

// Problem Constraints
// 1 <= A <= 10 pow 6

// Example Input
// Input 1:
// A = 4
// Input 2:
// A = 6

// Example Output
// Output 1:
// 0
// Output 2:
// 1

// Example Explanation

// Explanation 1:
// For A = 4, the sum of its proper divisors = 1 + 2 = 3, is not equal to 4.

// Explanation 2:
// For A = 6, the sum of its proper divisors = 1 + 2 + 3 = 6, is equal to 6.

package main

import (
	"fmt"
)

func isPerfectNumber(n int) int {
	sum := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			second := n / i
			if i != n {
				sum = sum + i
			}
			if i != second && second != n {
				sum = sum + second
			}
		}
	}
	if sum == n {
		return 1
	}
	return 0
}

func main() {
	n := 496
	output := isPerfectNumber(n)
	if output == 1 {
		fmt.Println(n, " is a Perfect Number.")
	} else {
		fmt.Println(n, " is not a Perfect Number.")
	}

}

// Time Complexity: O(N)
// Space Complexity: O(1)
