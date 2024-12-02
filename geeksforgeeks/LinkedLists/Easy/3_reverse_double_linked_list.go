package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
	prev *Node
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
		return
	}

	new_node.prev = l.tail
	l.tail.next = new_node
	l.tail = new_node
}

func printForward(l *List) {
	curr := l.head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}

func printBackward(l *List) {
	curr := l.tail
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.prev
	}
}
func reverseList(l *List) {
	cur := l.head
	var prev *Node
	var next *Node

	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}

	l.head = prev
	printForward(l)
}

func main() {
	l := &List{}
	l.add(2)
	l.add(12)
	l.add(121)
	reverseList(l)
}
