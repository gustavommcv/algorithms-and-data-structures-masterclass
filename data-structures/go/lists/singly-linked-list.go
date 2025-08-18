package lists

import (
	"errors"
	"fmt"
	"strings"
)

// Node represents a single element in a singly linked list.
// It contains a generic Value field to store data and a Next pointer
// to reference the subsequent node in the list.
type Node[T any] struct {
	Next  *Node[T] // Pointer to the next node in the list
	Value T        // The value stored in this node
}

// SinglyLinkedList represents a linear collection of elements where each element
// points to the next one in sequence. This implementation maintains references
// to both the Head (first element) and Tail (last element) of the list,
// along with a Length counter for efficient size tracking.
type SinglyLinkedList[T any] struct {
	Head   *Node[T] // Pointer to the first node in the list
	Tail   *Node[T] // Pointer to the last node in the list
	Length int      // Number of nodes in the list
}

// Push adds a new node with the given value to the end (tail) of the list.
// Returns the modified list to allow method chaining.
// Time complexity: O(1) as we maintain a Tail reference.
func (s *SinglyLinkedList[T]) Push(n T) *SinglyLinkedList[T] {
	newNode := &Node[T]{
		Next:  nil,
		Value: n,
	}

	if s.Head == nil {
		s.Head = newNode
		s.Tail = newNode
		s.Length++
		return s
	}

	s.Tail.Next = newNode
	s.Tail = newNode
	s.Length++
	return s
}

// Unshift adds a new node with the given value to the beginning (head) of the list.
// Returns the modified list to allow method chaining.
// Time complexity: O(1) as we only modify the Head reference.
func (s *SinglyLinkedList[T]) Unshift(n T) *SinglyLinkedList[T] {
	newNode := &Node[T]{
		Next:  nil,
		Value: n,
	}

	if s.Head == nil {
		s.Head = newNode
		s.Tail = newNode
		s.Length++
		return s
	}

	newNode.Next = s.Head
	s.Head = newNode
	s.Length++

	return s
}

// Get retrieves the node at the specified position in the list.
// Positions are zero-indexed (0 = head, length-1 = tail).
// Returns:
// - The node at the requested position, or
// - An error if:
//   - Position is negative
//   - Position is out of bounds (>= list length)
//   - List is empty
//
// Time complexity: O(n) in worst case as it may need to traverse the entire list.
func (s *SinglyLinkedList[T]) Get(position int) (*Node[T], error) {
	if position < 0 {
		return nil, errors.New("position must be positive")
	}

	if position >= s.Length {
		return nil, fmt.Errorf("position %d exceeds list length %d", position, s.Length)
	}

	current := s.Head

	if current == nil {
		return nil, errors.New("list is empty")
	}

	if position == 0 {
		return current, nil
	}

	if position == s.Length-1 {
		return s.Tail, nil
	}

	i := 0
	for {
		if i == position {
			return current, nil
		}

		current = current.Next

		if current == nil {
			return nil, nil
		}

		i++
	}
}

// Set updates the value of the node at the specified position in the list.
// Returns:
// - true if the value was successfully updated
// - false if:
//   - The position is out of bounds
//   - The list is empty
//   - The node couldn't be found
//
// Time complexity: O(n) in worst case as it uses Get() which may need to traverse the list.
func (s *SinglyLinkedList[T]) Set(index int, value T) bool {
	node, err := s.Get(index)

	if err != nil {
		return false
	}

	if node == nil {
		return false
	}

	node.Value = value

	return true
}

// Pop removes and returns the last node (tail) from the list.
// Handles special cases:
// - Returns nil if the list is empty
// - Properly resets Head and Tail when removing the last element
// Returns the removed node.
// Time complexity: O(n) as we need to find the penultimate node.
func (s *SinglyLinkedList[T]) Pop() *Node[T] {
	if s.Head == nil {
		return nil
	}

	if s.Head.Next == nil {
		currentHead := s.Head
		s.Head = nil
		s.Tail = nil
		s.Length--

		return currentHead
	}

	currentTail := s.Tail
	penultimateNode := s.penultimateNode()

	penultimateNode.Next = nil
	s.Tail = penultimateNode
	s.Length--

	return currentTail
}

// Shift removes and returns the first node (head) from the list.
// Returns nil if the list is empty.
// Time complexity: O(1) as we only modify the Head reference.
func (s *SinglyLinkedList[T]) Shift() *Node[T] {
	if s.Head == nil {
		return nil
	}

	currentHead := s.Head
	s.Head = s.Head.Next
	s.Length--

	return currentHead
}

// String returns a formatted string representation of the list showing:
// - All elements in square brackets, comma-separated
// - The current head node's value
// - The current tail node's value
// - The current length of the list
func (s *SinglyLinkedList[T]) String() string {
	var sb strings.Builder

	currentElement := s.Head
	if currentElement == nil {
		sb.WriteString(fmt.Sprintf("[] - Length: %v", s.Length))
		return sb.String()
	}

	sb.WriteString("[")
	for {
		if currentElement == nil {
			break
		}

		sb.WriteString(fmt.Sprintf("%v", currentElement.Value))

		if !(currentElement == s.Tail) {
			sb.WriteString(", ")
		}

		currentElement = currentElement.Next
	}
	sb.WriteString("]")

	sb.WriteString(fmt.Sprintf(" - Length: %v", s.Length))
	sb.WriteString(fmt.Sprintf(" / Head: %v / Tail: %v", s.Head.Value, s.Tail.Value))
	return sb.String()
}

// penultimateNode is a helper method that finds and returns the second-to-last node in the list.
// Returns nil if the list is empty or has only one element.
// Time complexity: O(n) as it must traverse the list until the penultimate node.
func (s *SinglyLinkedList[T]) penultimateNode() *Node[T] {
	currentElement := s.Head

	if currentElement == nil {
		return nil
	}

	for {
		if currentElement.Next == s.Tail {
			break
		}

		currentElement = currentElement.Next
	}

	return currentElement
}
