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

func delete_first_node(l *List) {
	if l.head == nil {
		return
	}
	head := l.head
	head_next := head.next
	head.next = nil
	l.head = head_next

	l.tail.next = l.head
}

func delete_last_node(l *List) {
	if l.head == nil {
		return
	}
	head := l.head
	curr := head.next
	var prev *Node
	for curr != nil && curr != head {
		if curr.data == l.tail.data {
			curr.next = nil
			l.tail = prev
			l.tail.next = head
		}
		prev = curr
		curr = curr.next
	}
}

func delete_node(l *List, node_val int) {
	if l.head == nil {
		return
	}
	if node_val == l.head.data {
		delete_first_node(l)
		return
	}

	if node_val == l.tail.data {
		delete_last_node(l)
		return
	}

	head := l.head
	curr := head.next
	prev := l.head
	for curr != nil && curr != head {
		if curr.data == node_val {
			prev.next = curr.next
			curr.next = nil
		}
		prev = curr
		curr = curr.next
	}
}

func main() {
	l := &List{}
	l.add(3)
	l.add(4)
	l.add(5)
	l.add(6)

	delete_node(l, 4)
	print(l)
}
