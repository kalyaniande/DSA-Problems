// Given a binary array nums, you should delete one element from it.

// Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.

// Example 1:

// Input: nums = [1,1,0,1]
// Output: 3
// Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.
// Example 2:

// Input: nums = [0,1,1,1,0,1,1,0,1]
// Output: 5
// Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].
// Example 3:

// Input: nums = [1,1,1]
// Output: 2
// Explanation: You must delete one element.

// Constraints:

// 1 <= nums.length <= 10pow5
// nums[i] is either 0 or 1.

package main

import (
	"fmt"
)

func longestSubarray(nums []int, k int) int {
	left := 0
	zero_count := 0
	max_length := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zero_count++
		}
		for zero_count > 1 {
			if nums[left] == 0 {
				zero_count--
			}
			left++
		}
		max_length = max(max_length, right-left)
	}
	return max_length
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	k := 2

	fmt.Println("Output: ", longestSubarray(nums, k))
}
