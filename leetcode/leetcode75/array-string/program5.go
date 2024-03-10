// Given a string s, reverse only all the vowels in the string and return it.

// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

// Example 1:

// Input: s = "hello"
// Output: "holle"
// Example 2:

// Input: s = "leetcode"
// Output: "leotcede"

// Constraints:

// 1 <= s.length <= 3 * 105
// s consist of printable ASCII characters.

package main

import (
	"fmt"
)

func reverseVowels(s string) string {
	output := []rune(s)
	for i, j := 0, len(s)-1; i < j; {
		i_char := string(s[i])
		j_char := string(s[j])
		if isVowel(i_char) && isVowel(j_char) {
			output[i], output[j] = output[j], output[i]
			i++
			j--
		} else {

			if isVowel(i_char) {
				j--
			} else {

				i++
			}
		}

	}
	return string(output)
}

func isVowel(s string) bool {
	return s == "a" || s == "e" || s == "i" || s == "o" || s == "u" || s == "A" || s == "E" || s == "I" || s == "O" || s == "U"
}

func main() {
	s := "hello"

	fmt.Println("Output: ", reverseVowels(s))
}
