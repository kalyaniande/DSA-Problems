// You are given an integer array A of size N.
// You have to perform B operations. In one operation, you can remove either the leftmost or the rightmost element of the array A.
// Find and return the maximum possible sum of the B elements that were removed after the B operations.

// NOTE: Suppose B = 3, and array A contains 10 elements, then you can:
// Remove 3 elements from front and 0 elements from the back, OR
// Remove 2 elements from front and 1 element from the back, OR
// Remove 1 element from front and 2 elements from the back, OR
// Remove 0 elements from front and 3 elements from the back.

// Problem Constraints
// 1 <= N <= 105
// 1 <= B <= N
// -10 pow 3 <= A[i] <= 10 pow 3

// Example Input
// Input 1:
// A = [5, -2, 3 , 1, 2]
//  B = 3

// Input 2:
//  A = [ 2, 3, -1, 4, 2, 1 ]
//  B = 4

// Example Output
// Output 1:
// 8

// Output 2:
// 9

// Example Explanation
// Explanation 1:
// Remove element 5 from front and element (1, 2) from back so we get 5 + 1 + 2 = 8

// Explanation 2:
// Remove the first element and the last 3 elements. So we get 2 + 4 + 2 + 1 = 9

package main

import (
	"fmt"
)

func solve(A []int, B int) int {
	sum := A[0]
	for i := 1; i < B; i++ {
		sum += A[i]
	}
	output := sum
	i := B - 1
	j := len(A) - 1
	for i >= 0 {
		sum -= A[i]
		sum += A[j]
		if sum > output {
			output = sum
		}
		i--
		j--
	}
	return output
}

func main() {
	A := []int{2, 3, -1, 4, 2, 1}
	B := 4
	fmt.Println("Output: ", solve(A, B))
}
