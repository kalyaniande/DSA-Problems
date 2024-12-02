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
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = new_node
}

func print(l *List) {
	curr := l.head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}

func remove_duplicates(l *List) {
	curr := l.head
	var prev *Node
	for curr != nil {
		if prev != nil && prev.data == curr.data {
			prev.next = curr.next
		} else {
			prev = curr
		}
		curr = curr.next
	}
	print(l)
}

func main() {
	l := &List{}
	l.add(2)
	l.add(2)
	l.add(4)
	l.add(5)
	remove_duplicates(l)
}
