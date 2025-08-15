package lists

import (
	"fmt"
	"strings"
)

// Contains a 'value' field as well as 'next' field
type node[T any] struct {
	next  *node[T]
	value T
}

// Singly linked list is a linear collection of data elements whose order is not given by their physical placement in memory.
type SinglyLinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

// Insert node at the end (tail)
func (s *SinglyLinkedList[T]) Push(n T) *SinglyLinkedList[T] {
	newNode := &node[T]{
		next:  nil,
		value: n,
	}

	if s.head == nil {
		s.head = newNode
		s.tail = newNode
		s.length++
		return s
	}

	s.tail.next = newNode
	s.tail = newNode
	s.length++
	return s
}

// Insert node at the beggining (head)
func (s *SinglyLinkedList[T]) Unshift(n T) *SinglyLinkedList[T] {
	newNode := &node[T]{
		next:  nil,
		value: n,
	}

	if s.head == nil {
		s.head = newNode
		s.tail = newNode
		s.length++
		return s
	}

	newNode.next = s.head
	s.head = newNode
	s.length++

	return s
}

// Remove a node at the end (tail)
func (s *SinglyLinkedList[T]) Pop() *SinglyLinkedList[T] {
	if s.head == nil {
		return nil
	}

	penultimateNode := s.penultimateNode()

	penultimateNode.next = nil
	s.tail = penultimateNode
	s.length--

	return s
}

func (s *SinglyLinkedList[T]) Shift() *SinglyLinkedList[T] {
	if s.head == nil {
		return nil
	}

	s.head = s.head.next
	s.length--

	return s
}

func (s *SinglyLinkedList[T]) penultimateNode() *node[T] {
	currentElement := s.head

	if currentElement == nil {
		return nil
	}

	for {
		if currentElement.next == s.tail {
			break
		}

		currentElement = currentElement.next
	}

	return currentElement
}

// To string method: [list] - head / tail / lenght
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

		sb.WriteString(fmt.Sprintf("%v", currentElement.value))

		if !(currentElement == s.tail) {
			sb.WriteString(", ")
		}

		currentElement = currentElement.next
	}
	sb.WriteString("]")

	sb.WriteString(fmt.Sprintf(" - Head: %v / Tail: %v / Length: %v", s.head.value, s.tail.value, s.length))
	return sb.String()
}
