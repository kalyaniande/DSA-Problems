// You are given an integer array nums consisting of n elements, and an integer k.

// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.

// Example 1:

// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75000
// Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
// Example 2:

// Input: nums = [5], k = 1
// Output: 5.00000

// Constraints:

// n == nums.length
// 1 <= k <= n <= 10pow5
// -10pow4 <= nums[i] <= 10pow4

package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	curr_sum := 0

	for i := 0; i < k; i++ {
		curr_sum += nums[i]
	}
	max_sum := curr_sum

	for i := k; i < len(nums); i++ {
		curr_sum = curr_sum - nums[i-k] + nums[i]
		if curr_sum > max_sum {
			max_sum = curr_sum
		}

	}

	return float64(float64(max_sum) / float64(k))
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4

	fmt.Println("Output: ", findMaxAverage(nums, k))
}
