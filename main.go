package main

import (
	"fmt"
)

type Node struct {
	size int
	data int
	next *Node
	head *Node
}

// constructor
func initialize() *Node {
	return &Node{}
}
func (node *Node) addFront(data int) {
	newNode := &Node{
		data: data,
	}
	// linked list is empty
	if node.head == nil {
		node.head = newNode

	} else {
		// add new head
		newNode.next = node.head
		node.head = newNode
	}

	node.size++

}

func (node *Node) addBack(data int) {
	newNode := &Node{
		data: data,
	}
	// linked list is empty
	if node.head == nil {
		node.head = newNode
	} else {
		// try to find last node
		current := node.head
		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}

	node.size++
}

func (node *Node) removeFirstElement() error {
	if node.head == nil {
		return fmt.Errorf("removeFront : List is empty")
	}
	node.head = node.head.next
	node.size--
	return nil
}

func (node *Node) removeLastElement() error {
	if node.head == nil {
		return fmt.Errorf("removeBack : List is empty")
	}
	// find previous node of last node
	var previous *Node
	current := node.head

	for current.next != nil {
		previous = current
		current = current.next
	}
	if previous != nil {
		previous.next = nil

	} else {
		node.head = nil
	}
	node.size--
	return nil
}

// get element of linked list
func (node *Node) getFirstElement() (int, error) {

	if node.head == nil {
		return -1, fmt.Errorf("Single List is empty")
	}

	return node.head.data, nil
}

// get last element of linked list
func (node *Node) getLastElement() (int, error) {
	if node.head == nil {
		return -1, fmt.Errorf("Single List is empty")
	}

	current := node.head

	for current.next != nil {
		current = current.next
	}

	return current.data, nil
}

func (node *Node) length() int {
	return node.size
}

// print whole list
func (node *Node) traverse() error {

	if node.head == nil {
		return fmt.Errorf("TranverseError: List is empty")
	}

	current := node.head

	for current != nil {
		if current.next == nil {
			fmt.Println(current.data)
		} else {
			fmt.Printf("%d -> ", current.data)
		}

		current = current.next
	}

	return nil
}
func main() {
	// example
	list := initialize()

	list.addBack(5)
	list.addBack(7)
	list.addBack(6)
	list.addBack(4)
	list.addBack(3)
	list.addBack(2)

	list.removeFirstElement()
	list.removeLastElement()

	list.traverse()

	fmt.Println(list.getFirstElement())
	fmt.Println(list.getLastElement())
	fmt.Println(list.length())

}
