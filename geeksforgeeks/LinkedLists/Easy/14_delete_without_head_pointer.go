package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func delete_node_without_head(node *Node) {
	if node == nil || node.next == nil {
		return
	}

	node.data = node.next.data
	node.next = node.next.next

	fmt.Println(node.data)
}

func create(data int) *Node {
	new_node := &Node{data: data}
	new_node.next = nil
	return new_node
}

func print(head *Node) {
	temp := head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func main() {
	head := &Node{}

	head = create(5)
	head.next = create(4)
	head.next.next = create(3)
	delete := create(2)
	head.next.next.next = delete
	head.next.next.next.next = create(1)

	// print(head)

	// del := head.next.next.next
	delete_node_without_head(delete)
	print(head)
}
