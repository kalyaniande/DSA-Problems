// Say you have an array, A, for which the ith element is the price of a given stock on day i.
// If you were only permitted to complete at most one transaction (ie, buy one and sell one share of the stock),
// design an algorithm to find the maximum profit.
// Return the maximum possible profit.

// Problem Constraints
// 0 <= A.size() <= 700000
// 1 <= A[i] <= 10 pow 7

// Example Input
// Input 1:
// A = [1, 2]

// Input 2:
// A = [1, 4, 5, 2, 4]

// Example Output
// Output 1:
// 1

// Output 2:
// 4

// Example Explanation
// Explanation 1:
// Buy the stock on day 0, and sell it on day 1.

// Explanation 2:
// Buy the stock on day 0, and sell it on day 2.

package main

import (
	"fmt"
)

func maxProfit(A []int) int {

	min := A[0]
	profit := 0

	for i := 0; i < len(A); i++ {
		if A[i]-min > profit {
			profit = A[i] - min
		}

		if A[i] < min {
			min = A[i]
		}
	}
	return profit
}

func main() {
	A := []int{1, 4, 5, 2, 4}
	fmt.Println("Max Profit: ", maxProfit(A))
}
