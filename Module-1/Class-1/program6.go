// Given two integers A and B. A represents the count of mangoes and B represent the count of slices of mangoes. Mango can be cut into three slices of mangoes. A glass of mango shake can be formed by two slices of mangoes.

// Find the maximum number of glasses of mango shakes you can make with A mangoes and B slices of mangoes initially.

// Constraints
// 0 <= A, B <= 10^8

// For Example

// Input 1:
//     A = 19
//     B = 0
// Output 1:
//     28

// Input 2:
//     A = 7
//     B = 1
// Output 2:
//     11
package main

import (
	"fmt"
)

func solve(A int, B int) int {
	mango_slices := A * 3
	total_slices := mango_slices + B
	return int(total_slices / 2)
}

func main() {
	A := 19
	B := 2
	fmt.Println("Output:", solve(A, B))
}
