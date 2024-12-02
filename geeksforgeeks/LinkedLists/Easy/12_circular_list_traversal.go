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
	tail *Node
}

func (l *List) add(data int) {
	new_node := &Node{data: data}
	if l.head == nil {
		l.head = new_node
		l.tail = new_node
		new_node.next = l.head
		return
	}

	l.tail.next = new_node
	l.tail = new_node
	new_node.next = l.head

}

func print(l *List) {
	head := l.head
	if head == nil {
		return
	}

	fmt.Println(head.data)
	curr := head.next
	for curr != head {
		fmt.Println(curr.data)
		curr = curr.next
	}
}

func main() {
	l := &List{}
	l.add(3)
	l.add(4)
	l.add(5)
	l.add(6)
	print(l)
}
