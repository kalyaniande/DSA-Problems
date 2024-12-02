package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) add(data int) {
	new_node := &Node{data: data}

	if l.head == nil {
		l.head = new_node
	}

}

func main() {
	fmt.Println("Hi")
}
