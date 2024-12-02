package main

import (
	"fmt"
)

func fact(n int) int {
	if n <= 1 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	n := 10
	fact := fact(n)
	fmt.Println("factorial of", n, " is ", fact)
}
