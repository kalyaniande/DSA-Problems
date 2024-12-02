package main

import (
	"fmt"
)

func print(n int) {
	if n >= 1 {
		fmt.Println(n)
		print(n - 1)
	}

	return
}

func main() {
	n := 10
	print(n)
}
