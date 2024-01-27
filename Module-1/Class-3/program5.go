// Given an array of integers A, find and return the product array of the same size
// where the ith element of the product array will be equal to the product of all the elements divided by the ith element of the array.

// Note: It is always possible to form the product array with integer (32 bit) values.
// Solve it without using the division operator.

// Problem Constraints
// 2 <= length of the array <= 1000
// 1 <= A[i] <= 10

// Input 1:
//     A = [1, 2, 3, 4, 5]
// Output 1:
//     [120, 60, 40, 30, 24]

// Input 2:
//     A = [5, 1, 10, 1]
// Output 2:
//     [10, 50, 5, 50]

package main

import (
	"fmt"
)

func solve(A []int) []int {
	product := 1
	for i := 0; i < len(A); i++ {
		product = product * A[i]
	}

	for i := 0; i < len(A); i++ {
		A[i] = product / A[i]
	}
	return A
}

func main() {
	A := []int{1, 2, 3, 4, 5}
	fmt.Println("Output:", solve(A))
}

// pref := make([]int, len(A))
// 	suf := make([]int, len(A))
// 	output := make([]int, len(A))
// 	pref[0] = A[0]
// 	suf[len(A)-1] = 1
// 	for i := 1; i < len(A); i++ {
// 		pref[i] = pref[i-1] * A[i]
// 		suf[len(A)-i-1] = suf[len(A)-i] * A[len(A)-i]
// 	}
// 	for i := 0; i < len(A); i++ {
// 		if i == 0 {
// 			output[i] = suf[i]
// 		} else if i == (len(A) - 1) {
// 			output[i] = pref[i-1]

// 		} else {

// 			output[i] = pref[i-1] * suf[i]
// 		}
// 	}
// return output
