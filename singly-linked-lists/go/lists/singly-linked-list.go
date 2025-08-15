package lists

import (
	"fmt"
	"strings"
)

type node[T any] struct {
	next *node[T]
	data T
}

type SinglyLinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func (s *SinglyLinkedList[T]) Append(n T) {
	newNode := &node[T]{
		next: nil,
		data: n,
	}

	if s.head == nil {
		s.head = newNode
		s.length++
		return
	}

	if s.length == 1 {
		tail := newNode
		s.head.next = tail
		s.tail = tail
		s.length++
		return
	}

	if s.tail != nil {
		s.tail.next = newNode
		s.tail = newNode
	}
}

func (s *SinglyLinkedList[T]) String() string {
	var sb strings.Builder

	currentElement := s.head
	if currentElement == nil {
		return ""
	}

	sb.WriteString("[")
	for {
		if currentElement == nil {
			break
		}

		sb.WriteString(fmt.Sprintf("%v,", currentElement.data))
		currentElement = currentElement.next
	}
	sb.WriteString("]")

	return sb.String()
}
