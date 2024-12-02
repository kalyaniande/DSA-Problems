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

func merge_list_at_alternate_positions(l1 *List, l2 *List) {
	if l1.head == nil || l2.head == nil {
		return
	}

	curr1 := l1.head
	curr2 := l2.head
	for curr1 != nil && curr2 != nil {
		// prev := curr1
		// next := curr1.next
		// l2_node := &Node{data: curr2.data}
		// prev.next = l2_node
		// prev.next.next = next
		// curr2 = curr2.next
		// curr1 = next

		l1_next := curr1.next
		l2_next := curr2.next

		curr2.next = curr1.next
		curr1.next = curr2

		curr1 = l1_next
		curr2 = l2_next
	}
	print(l1)

}

func main() {
	l1 := &List{}
	l1.add(1)
	l1.add(2)
	l1.add(3)
	l1.add(4)
	l1.add(5)

	l2 := &List{}
	l2.add(6)
	l2.add(7)
	l2.add(8)
	l2.add(9)
	l2.add(10)
	l2.add(11)
	l2.add(12)
	merge_list_at_alternate_positions(l1, l2)
}
