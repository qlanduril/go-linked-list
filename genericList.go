package main

import (
	"fmt"
)

type GenericNode[E any] struct {
	data E
	next *GenericNode[E]
}

func NewGenericNode[E any](value E) *GenericNode[E] {
	return &GenericNode[E]{
		data: value,
		next: nil,
	}

}

func (n *GenericNode[E]) printGenericNode() {

	fmt.Println("Node address= ", &n, "\ndata= ", n.data, "\nnext address= ", &n.next)
}

type GenericList[T any] struct {
	head *GenericNode[T]
}

func NewGenericList[T any]() *GenericList[T] {
	return &GenericList[T]{
		head: nil,
	}
}

func (l *GenericList[T]) AddItem(value T) {
	tmp := NewGenericNode[T](value)

	tmp.next = l.head
	l.head = tmp
}

func (l *GenericList[E]) PrintList() {
	printGenericList(l.head)
}

func printGenericList[E any](head *GenericNode[E]) {
	if head == nil {
		return
	}

	head.printGenericNode()
	printGenericList(head.next)
}
