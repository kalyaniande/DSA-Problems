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

func delete_last_occurance(l *List, element int) {
	curr := l.head
	var last_occurance *Node
	var last_occurance_prev *Node
	var prev *Node
	for curr != nil {
		if curr.data == element {
			last_occurance = curr
			last_occurance_prev = prev
		}
		prev = curr
		curr = curr.next
	}

	last_occurance_prev.next = last_occurance.next
	last_occurance.next = nil
	print(l)
}

func main() {
	l := &List{}
	l.add(9)
	l.add(45)
	l.add(9)
	l.add(45)
	l.add(9)
	delete_last_occurance(l, 45)
}
