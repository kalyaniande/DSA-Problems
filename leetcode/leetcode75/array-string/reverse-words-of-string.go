// Given an input string s, reverse the order of the words.

// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

// Return a string of the words in reverse order concatenated by a single space.

// Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

// Example 1:

// Input: s = "the sky is blue"
// Output: "blue is sky the"
// Example 2:

// Input: s = "  hello world  "
// Output: "world hello"
// Explanation: Your reversed string should not contain leading or trailing spaces.
// Example 3:

// Input: s = "a good   example"
// Output: "example good a"
// Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

// Constraints:

// 1 <= s.length <= 104
// s contains English letters (upper-case and lower-case), digits, and spaces ' '.
// There is at least one word in s.

// Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?

package main

import (
	"fmt"
)

func reverseWords(s string) string {
	var output string
	word_start_index := -1

	prev_char := string(s[0])

	for i := 0; i < len(s); i++ {

		if string(s[i]) == " " {
			if prev_char == " " {
				word_start_index++
				continue
			}
			word := s[word_start_index+1 : i]
			if output == "" {
				output = word + output
			} else {
				output = word + " " + output
			}

			word_start_index = i
		}
		prev_char = string(s[i])
	}
	if word_start_index != len(s)-1 {
		if output == "" {
			output = s[word_start_index+1:] + output
		} else {
			output = s[word_start_index+1:] + " " + output
		}
	}
	return output
}

func main() {
	s := "the sky is blue"

	fmt.Println("Output: ", reverseWords(s))
}
