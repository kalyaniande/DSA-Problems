// Given a number n. Return 1 if n is prime and return 0 if not.

// Note :
// The value of n can cross the range of Integer.

// Problem Constraints
// 1 <= n <= 10 pow 9
package main

import (
	"fmt"
)

func isPrime(n int) bool {

	if n == 1 {
		return false
	}

	for i := 2; i*i < n; i++ {
		if i != n && n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	n := 6

	is_prime := isPrime(n)
	if is_prime {
		fmt.Println(n, " is a Prime Number")
	} else {
		fmt.Println(n, " is not a Prime Number")
	}
}
