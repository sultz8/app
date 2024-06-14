package main

import "fmt"

type Node struct {
    value interface{}
	next *Node
	prev *Node
}

type List struct {
    head *Node
	tail *Node
	size int
}

func (l *List) PushFront(val interface{}) (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = &InsertListError{msg: "Insertion failed."}
		}
	}()
	
	newNode := &Node{
		value: val,
		next: nil,
		prev: nil,
	}
	
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
	l.size = l.size + 1
	
	return nil
}

type ListBaseError struct {
    msg string
}

type InsertListError ListBaseError

type EmptyListError ListBaseError

func (e *InsertListError) Error() string {
	return e.msg
}

func (e *EmptyListError) Error() string {
	return e.msg
}

func (l *List) Print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Print(currentNode.value, " ")
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (l *List) popBack() (node *Node, err error) {
	var lastNode *Node = nil
	if l.head == nil {
		return nil, &EmptyListError{msg: "list is empty!"}
	}

	if l.size == 1 {
		lastNode = l.head
		l.head = nil
		l.tail = nil
		l.size = l.size - 1

		return lastNode, nil
	}

	lastNode = l.tail
	l.tail = l.tail.prev
	l.tail.next = nil
	lastNode.prev = nil
	l.size = l.size - 1

	return lastNode, nil
}