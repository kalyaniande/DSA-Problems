// You are given a constant array A.
// You are required to return another array which is the reversed form of the input array.

// Problem Constraints
// 1 <= A.size() <= 10000
// 1 <= A[i] <= 10000

// Example Input
// Input 1:
// A = [1,2,3,2,1]

// Input 2:
// A = [1,1,10]

// Example Output
// Output 1:
//  [1,2,3,2,1]

// Output 2:
//  [10,1,1]

// Example Explanation
// Explanation 1:
// Reversed form of input array is same as original array

// Explanation 2:
// Clearly, Reverse of [1,1,10] is [10,1,1]

package main

import (
	"fmt"
)

func reverseArray(A []int) []int {

	i := 0
	j := len(A) - 1
	for i <= j {
		A[i], A[j] = A[j], A[i]
		i++
		j--
	}
	return A
}

func main() {
	A := []int{1, 2, 3, 4}
	fmt.Println("Output:", reverseArray(A))
}
