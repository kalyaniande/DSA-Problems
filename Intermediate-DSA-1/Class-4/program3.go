// Given an array A, find the size of the smallest subarray
// such that it contains at least one occurrence of the maximum value of the array
// and at least one occurrence of the minimum value of the array.

// Problem Constraints
// 1 <= |A| <= 2000

// Example Input
// Input 1:
// A = [1, 3, 2]

// Input 2:
// A = [2, 6, 1, 6, 9]

// Example Output
// Output 1:
// 2

// Output 2:
// 3

// Example Explanation
// Explanation 1:
// Take the 1st and 2nd elements as they are the minimum and maximum elements respectievly.

// Explanation 2:
// Take the last 3 elements of the array.

// Observations:
// There is only one min and one max in the subarray
// min and max are always present in the corners

package main

import (
	"fmt"
)

func solve(A []int) int {

	min := A[0]
	max := A[0]

	for i := 0; i < len(A); i++ {
		if A[i] > max {
			max = A[i]
		}

		if A[i] < min {
			min = A[i]
		}
	}
	if min == max {
		return 1
	}
	si := -1
	ei := -1
	output := len(A)
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] == min {
			si = i
			if ei != -1 {
				len := ei - si + 1
				if len < output {
					output = len
				}
			}
		} else if A[i] == max {
			ei = i
			if si != -1 {
				len := si - ei + 1
				if len < output {
					output = len
				}
			}
		}
	}
	return output
}

func main() {
	A := []int{2, 6, 1, 6, 9}
	fmt.Println("Output: ", solve(A))
}
