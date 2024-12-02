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

func printList(l *List) {
	curr := l.head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
	fmt.Println()
}

func (l *List) add(value int) {
	newNode := &Node{data: value}
	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func getMiddle(l *List) {
	curr := l.head
	if curr == nil {
		return
	}
	length := 1
	for curr.next != nil {
		length++
		curr = curr.next
	}
	mid_index := (length / 2) + 1
	curr = l.head
	cur_index := 1
	for cur_index < mid_index {
		cur_index++
		curr = curr.next
	}
	fmt.Println("Middle Node data: ", curr.data)
}

func main() {
	list := &List{}
	list.add(3)
	list.add(5)
	list.add(6)
	list.add(9)
	// list.add(2)

	printList(list)
	getMiddle(list)
}
