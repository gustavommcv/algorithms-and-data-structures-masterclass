package main

import (
	"fmt"

	"github.com/gustavommcv/algorithms-and-data-structures-masterclass/singly-linked-lists/lists"
)

func main() {
	list := new(lists.SinglyLinkedList[int])

	list.Push(0)
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)

	list.Set(0, 12)
	list.Insert(3, 25)
	list.Remove(3)
	list.Remove(0)

	fmt.Println("list: " + list.String())
}
