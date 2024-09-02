package main

import (
	"fmt"

	ng "github.com/qlanduril/go-linked-list/nongeneric"
)

func main() {
	fmt.Println("hellope")

	list := ng.NewList()

	list.AddItem(12)

	for i := 85; i < 93; i++ {
		list.AddItem(i)
	}

	list.PrintList()
}
