package main

import (
	"fmt"

	"github.com/gustavommcv/algorithms-and-data-structures-masterclass/singly-linked-lists/lists"
)

func main() {
	l := lists.SinglyLinkedList[int]{}

	l.Append(4)
	l.Append(4)
	l.Append(5)
	l.Append(4)
	l.Append(8)

	fmt.Println("list: " + l.String())
}
