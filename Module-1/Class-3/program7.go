// Given an array of integers A
// Find the number of special indices in the given array
// An index is said to be a special index, if after deleting it suma all even indices equal to sum of all odd indices.

// Problem Constraints
// 2 <= length of the array <= 1000
// 1 <= A[i] <= 10

// Input 1:
//   A = [4, 3, 2, 7, 6, -2]

// Output 1:
//     2

package main

import (
	"fmt"
)

func solve(A []int) int {
	count := 0
	pfeven := make([]int, len(A))
	pfodd := make([]int, len(A))
	pfeven[0] = A[0]
	pfodd[0] = 0
	for i := 1; i < len(A); i++ {
		if i%2 == 0 {
			pfeven[i] = pfeven[i-1] + A[i]
			pfodd[i] = pfodd[i-1]
		} else {
			pfodd[i] = pfodd[i-1] + A[i]
			pfeven[i] = pfeven[i-1]
		}
	}
	for i := 0; i < len(A); i++ {
		if i == 0 {
			even_sum := pfodd[len(A)-1] - pfodd[i]
			odd_sum := pfeven[len(A)-1] - pfeven[i]

			if even_sum == odd_sum {
				count++
			}
		} else {
			even_sum := pfeven[i-1] + pfodd[len(A)-1] - pfodd[i]
			odd_sum := pfodd[i-1] + pfeven[len(A)-1] - pfeven[i]

			if even_sum == odd_sum {
				count++
			}
		}

	}
	return count
}

func main() {
	A := []int{4, 3, 2, 7, 6, -2}
	fmt.Println("Output:", solve(A))
}
