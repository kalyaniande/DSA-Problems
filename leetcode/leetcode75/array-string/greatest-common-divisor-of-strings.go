// For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

// Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

// Example 1:

// Input: str1 = "ABCABC", str2 = "ABC"
// Output: "ABC"
// Example 2:

// Input: str1 = "ABABAB", str2 = "ABAB"
// Output: "AB"
// Example 3:

// Input: str1 = "LEET", str2 = "CODE"
// Output: ""

// Constraints:

// 1 <= str1.length, str2.length <= 1000
// str1 and str2 consist of English uppercase letters.

package main

import (
	"fmt"
)

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	s1_length := len(str1)
	s2_length := len(str2)

	gcd := s1_length
	if s2_length < s1_length {
		gcd = s2_length
	}
	for gcd > 0 {
		if s1_length%gcd == 0 && s2_length%gcd == 0 {
			break
		}
		gcd--
	}

	return str1[:gcd]
}
func main() {
	word1 := "ABCABC"
	word2 := "ABC"
	fmt.Println("Output: ", gcdOfStrings(word1, word2))
}

// Time Complexity: O(N)
// Space Complexity: O(1)
