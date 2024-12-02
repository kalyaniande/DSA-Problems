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

func delete_middle(l *List) {
	curr := l.head
	count := 0
	for curr != nil {
		count++
		curr = curr.next
	}

	fmt.Println("count: ", count)
	mid_index := (count / 2) + 1
	fmt.Println("mid_index: ", mid_index)

	var mid_node *Node
	var prev *Node
	var mid_prev *Node
	curr = l.head
	index := 1
	for curr != nil {
		if index == mid_index {
			mid_node = curr
			mid_prev = prev
		}
		prev = curr
		curr = curr.next
		index++
	}
	mid_prev.next = mid_node.next
	mid_node.next = nil
	print(l)
}

func main() {
	l := &List{}
	l.add(9)
	l.add(3)
	l.add(95)
	l.add(45)
	l.add(23)
	delete_middle(l)
}
