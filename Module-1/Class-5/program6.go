// Given an array of integers A, a subarray of an array is said to be good if it fulfills any one of the criteria:
// 1. Length of the subarray is be even, and the sum of all the elements of the subarray must be less than B.
// 2. Length of the subarray is be odd, and the sum of all the elements of the subarray must be greater than B.
// Your task is to find the count of good subarrays in A.

// Problem Constraints
// 1 <= len(A) <= 5 x 10 pow 3
// 1 <= A[i] <= 10 pow 3
// 1 <= B <= 10 pow 7

// Example Input
// Input 1:
// A = [1, 2, 3, 4, 5]
// B = 4

// Input 2:
// A = [0, 0, 0, 1, 1, 0, 1]
//  B = 0

// Example Output
// Output 1:
// 6

// Output 2:
// 36

// Example Explanation
// Explanation 1:
// Even length good subarrays = {1, 2}
// Odd length good subarrays = {1, 2, 3}, {1, 2, 3, 4, 5}, {2, 3, 4}, {3, 4, 5}, {5}

// Explanation 2:
// There are 36 good subarrays

package main

import (
	"fmt"
)

func solve(A []int, B int) int {
	output := 0
	length := len(A)

	for i := 0; i < length; i++ {
		sum := 0
		for j := i; j < length; j++ {
			sum += A[j]
			if i == j {
				if sum > B {
					output++
				}
			} else {
				if (j-i+1)%2 == 0 {
					if sum < B {
						output++
					}
				} else {
					if sum > B {
						output++
					}
				}
			}
		}

	}

	return output
}

func main() {
	A := []int{13, 16, 16, 15, 9, 16, 2, 7, 6, 17, 3, 9}
	B := 65
	fmt.Println("Output: ", solve(A, B))
}
