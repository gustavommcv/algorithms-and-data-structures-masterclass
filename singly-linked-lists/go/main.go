package main

import (
	"fmt"

	"github.com/gustavommcv/algorithms-and-data-structures-masterclass/singly-linked-lists/lists"
)

func main() {
	list := new(lists.SinglyLinkedList[int])

	list.Push(0)
	list.Push(5)
	list.Push(4)
	list.Push(8)
	list.Push(4).Push(10).Push(6)

	fmt.Println("list: " + list.String())

	fmt.Println("list: " + list.Pop().Pop().String())

	fmt.Println("list: " + list.Shift().String())
}
