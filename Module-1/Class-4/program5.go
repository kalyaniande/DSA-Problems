// You are given a string S, and you have to find all the amazing substrings of S.
// An amazing Substring is one that starts with a vowel (a, e, i, o, u, A, E, I, O, U).

// Problem Constraints
// 1 <= length(S) <= 1e6
// S can have special characters

// Example Explanation
// Input
// ABEC

// Output
//     6

// Explanation
//     Amazing substrings of given string are :
//     1. A
//     2. AB
//     3. ABE
//     4. ABEC
//     5. E
//     6. EC
//     here number of substrings are 6 and 6 % 10003 = 6.

package main

import (
	"fmt"
	"strings"
)

func solve(A string) int {

	output := 0
	length := len(A)
	for i, ch := range A {
		lower_character := strings.ToLower(string(ch))

		if lower_character == "a" || lower_character == "e" || lower_character == "i" || lower_character == "o" || lower_character == "u" {
			output += +length - i
		}
	}
	return output
}

func main() {
	A := "ABEC"
	fmt.Println("Output: ", solve(A))
}
