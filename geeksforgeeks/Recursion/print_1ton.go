package main

import (
	"fmt"
)

func print(n int) {
	if n >= 1 {
		print(n - 1)
		fmt.Println(n)
	}

	return
}

func main() {
	n := 10
	print(n)
}
