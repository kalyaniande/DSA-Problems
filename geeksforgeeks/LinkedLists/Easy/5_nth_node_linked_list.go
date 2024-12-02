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

func printList(l *List) {
	curr := l.head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
	fmt.Println()
}

func findNthNode(l *List, n int) {
	curr := l.head
	count := 0
	for curr != nil {
		count++
		curr = curr.next
	}
	fmt.Println("count:", count)
	if n > 5 {
		fmt.Println("nth element not found", -1)
	}

	curr = l.head
	index := 1
	for curr != nil {
		if index == n {
			fmt.Println("nth element:", curr.data)
			return
		}
		curr = curr.next
		index++
	}
}

func main() {
	l := &List{}
	l.add(2)
	l.add(98)
	l.add(34)
	l.add(56)
	l.add(8)

	n := 5
	findNthNode(l, n)
}
