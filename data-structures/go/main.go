package main

import (
	"fmt"

	"github.com/gustavommcv/algorithms-and-data-structures-masterclass/singly-linked-lists/lists"
)

func main() {
	list := new(lists.SinglyLinkedList[int])

	list.Push(0)
	list.Push(0)

	list.Shift()
	list.Shift()

	fmt.Println("list: " + list.String())
}
