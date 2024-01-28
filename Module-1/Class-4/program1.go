// You have given a string A having Uppercase English letters.
// You have to find how many times subsequence "AG" is there in the given string.
// NOTE: Return the answer modulo 109 + 7 as the answer can be very large.

// Problem Constraints
// 1 <= length(A) <= 10 pow 5

// Example Input
// Input 1:
// A = "ABCGAG"

// Input 2:
//  A = "GAB"

// Example Output
// Output 1:
// 3

// Output 2:
// 0

// Example Explanation
// Explanation 1:
// Subsequence "AG" is 3 times in given string

// Explanation 2:
// There is no subsequence "AG" in the given string.

package main

import (
	"fmt"
	"strings"
)

func solve(s string) int {
	output := 0
	arr := strings.Split(s, "")

	// Left to Right Approach
	letter_count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == "A" {
			letter_count++
		} else if arr[i] == "G" {
			output = output + letter_count
		}
	}

	// Right to Left Approach
	letter_count = 0
	output = 0
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == "G" {
			letter_count++
		} else if arr[i] == "A" {

			output = output + letter_count
		}
	}
	return output
}

func main() {
	A := "ABCGAG"
	fmt.Println("Output: ", solve(A))
}
