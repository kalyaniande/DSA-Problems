package main

import (
	"fmt"
)

func getNumberOfIslands(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}

	return count
}

func main() {
	grid := [][]int{[]int{1, 1, 0, 0, 0}, []int{0, 1, 0, 0, 0}, []int{1, 0, 0, 1, 1}, []int{0, 0, 0, 0, 0}, []int{1, 0, 1, 0, 0}}

	fmt.Println("number of islands:", getNumberOfIslands(grid))

	grid2 := [][]int{[]int{1, 1, 0, 0, 0}, []int{1, 1, 0, 0, 0}, []int{0, 0, 1, 0, 0}, []int{0, 0, 0, 1, 1}}

	fmt.Println("number of islands:", getNumberOfIslands(grid2))
}
