// A wire connects N light bulbs.
// Each bulb has a switch associated with it; however, due to faulty wiring, a switch also changes the state of all the bulbs to the right of the current bulb.
// Given an initial state of all bulbs, find the minimum number of switches you have to press to turn on all the bulbs.
// You can press the same switch multiple times.
// Note: 0 represents the bulb is off and 1 represents the bulb is on.

// Problem Constraints
// 0 <= N <= 5×10 pow 5
// 0 <= A[i] <= 1

// Example Input
// Input 1:
// A = [0, 1, 0, 1]

// Input 2:
//  A = [1, 1, 1, 1]

// Example Output
// Output 1:
// 4

// Output 2:
// 0

// Example Explanation
// Explanation 1:
// press switch 0 : [1 0 1 0]
//  press switch 1 : [1 1 0 1]
//  press switch 2 : [1 1 1 0]
//  press switch 3 : [1 1 1 1]

// Explanation 2:
// There is no need to turn any switches as all the bulbs are already on.

package main

import (
	"fmt"
)

func solve(A []int) int {
	no_of_switches := 0
	length := len(A)
	if length == 0 {
		return 0
	}
	if A[0] == 0 {
		no_of_switches++
	}
	for i := 1; i < length; i++ {
		if A[i-1] != A[i] {
			no_of_switches++
		}
	}
	return no_of_switches
}

func main() {
	A := []int{0, 1, 0, 1}
	fmt.Println("Output: ", solve(A))
}
