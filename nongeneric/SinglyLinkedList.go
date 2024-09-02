package nongeneric

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func NewNode(n int) *Node {
	return &Node{
		data: n,
		next: nil,
	}
}

func (n *Node) PrintNode() {
	fmt.Println("Node address= ", &n, "\ndata= ", n.data, "\nnext address= ", &n.next)
}

// linked list type
// grow at head
type List struct {
	head *Node
}

func NewList() *List {
	return &List{
		head: nil,
	}
}

func (l *List) AddItem(n int) {

	tmp := NewNode(n)

	// nil head instert
	if l.head == nil {
		l.head = tmp
		return
	}

	// add to front
	tmp.next = l.head
	l.head = tmp
}

func (l *List) PrintList() {
	if l.head == nil {
		return
	}
	l.head.PrintNode()
	printList(l.head.next)
}

func printList(h *Node) {
	if h == nil {
		return
	}
	h.PrintNode()
	printList(h.next)
}
