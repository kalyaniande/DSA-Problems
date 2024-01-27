package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func count() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		i := 0
		for ; i < 1e6; i++ {
		}
		_ = i
	}
}

func main() {
	a := [...]int{12, 78, 50}
	fmt.Println(a[0])

	// reference types:
	// Slice
	// Maps
	// Channels

	for i, v := range a {
		fmt.Println(i, v)
	}

	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	fruitslice = fruitarray[:cap(fruitslice)]                                        //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
	fmt.Println(fruitslice)

	// sl := make([]int, 4)
	// mapex := make(map[string]int)

	s := "Señor"

	for index, v := range s {
		fmt.Printf("%c starts at byte %d\n", v, index)
	}

}
