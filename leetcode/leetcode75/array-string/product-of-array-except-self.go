// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
// Example 2:

// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]

// Constraints:

// 2 <= nums.length <= 105
// -30 <= nums[i] <= 30
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)

package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	length := len(nums)
	output := make([]int, length)
	output[0], output[length-1] = 1, 1

	// Calculate the prefix product array
	for i := 1; i < len(nums); i++ {
		output[i] = output[i-1] * nums[i-1]
	}

	// calculate the suffix product and multiply  with prefix product to get the result
	suffix_product := 1
	for i := length - 2; i >= 0; i-- {
		suffix_product *= nums[i+1]
		output[i] = output[i] * suffix_product
	}

	return output
}
func main() {
	nums := []int{10, 3, 5, 6, 2}

	fmt.Println("Output: ", productExceptSelf(nums))
}
