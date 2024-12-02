// You are given an integer array nums and an integer k.

// In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

// Return the maximum number of operations you can perform on the array.

// Example 1:

// Input: nums = [1,2,3,4], k = 5
// Output: 2
// Explanation: Starting with nums = [1,2,3,4]:
// - Remove numbers 1 and 4, then nums = [2,3]
// - Remove numbers 2 and 3, then nums = []
// There are no more pairs that sum up to 5, hence a total of 2 operations.
// Example 2:

// Input: nums = [3,1,3,4,3], k = 6
// Output: 1
// Explanation: Starting with nums = [3,1,3,4,3]:
// - Remove the first two 3's, then nums = [1,4,3]
// There are no more pairs that sum up to 6, hence a total of 1 operation.

// Constraints:

// 1 <= nums.length <= 10 pow 5
// 1 <= nums[i] <= 10 pow 9
// 1 <= k <= 10 pow 9

package main

import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	result := 0
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			result++
			left++
			right--
		} else if sum < k {
			left++
		} else {
			right--
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	k := 5

	fmt.Println("Output: ", maxOperations(nums, k))
}

// Time Complexity: O(N)
// Space Complexity: O(1)
