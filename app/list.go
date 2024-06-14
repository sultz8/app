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

type InsertListError struct {
    msg string
}

func (e *InsertListError) Error() string {
	return e.msg
}

func (l *List) Print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Print(currentNode.value, " ")
		currentNode = currentNode.next
	}
}