package main

import (
	"fmt"
)

type Node struct {
	next *Node
	data int
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
}

func (l *List) add(val int) {
	newNode := &Node{data: val}
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
	printList(l)
}

func main() {
	l := &List{}
	l.add(2)
	l.add(3)
	l.add(4)
	l.add(5)
	l.add(6)
	fmt.Println("Reverse Linked List: ")
	reverseList(l)
}
