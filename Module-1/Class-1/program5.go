// You will be given an integer n. You need to return the count of prime numbers less than or equal to n.

// Problem Constraints
// 0 <= n <= 10^3

// Example Input
// Input 1:
// 19
// Input 2:
// 1

// Example Output
// Output 1:
// 8
// Output 2:
// 0

package main

import (
	"fmt"
)

func isPrime(n int) int {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 && i != 1 && i != n {
			return 0
		}
	}
	return 1
}

func primeNumbersCount(n int) int {
	count := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) == 1 {
			count++
		}
	}
	return count
}

func main() {
	n := 19
	fmt.Println("Count of prime numbers less than or equal to ", n, ":", primeNumbersCount(n))
}
