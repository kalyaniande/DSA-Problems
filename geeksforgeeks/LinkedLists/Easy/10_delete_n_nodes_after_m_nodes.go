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

func delete_n_nodes_after_m_nodes(l *List, n int, m int) {
	curr := l.head

	for curr != nil {
		// Skip m nodes
		for i := 1; i < m && curr != nil; i++ {
			curr = curr.next
		}

		// If we reached end of list, then return
		if curr == nil {
			return
		}
		// fmt.Println(curr.data)
		// Start from next node and delete n nodes
		next_node := curr.next
		for i := 1; i <= n && next_node != nil; i++ {
			temp := next_node
			next_node = next_node.next
			temp.next = nil
		}
		//Link the previous list with remaining nodes
		curr.next = next_node

		// Set current pointer for next iteration
		curr = next_node
	}
	print(l)
}

func main() {
	l := &List{}
	l.add(1)
	l.add(2)
	l.add(3)
	l.add(4)
	l.add(5)
	l.add(6)

	n := 2
	m := 2
	delete_n_nodes_after_m_nodes(l, n, m)
}
