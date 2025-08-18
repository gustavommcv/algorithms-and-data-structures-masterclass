package lists

import (
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
		// First node becomes both Head and Tail
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
		// First node becomes both Head and Tail
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
// - (*Node, true) if the node was found
// - (nil, false) if:
//   - Position is negative
//   - Position is out of bounds (>= list length)
//   - List is empty
//
// Time complexity: O(n) in worst case as it may need to traverse the entire list.
func (s *SinglyLinkedList[T]) Get(position int) (*Node[T], bool) {
	if position < 0 || position >= s.Length {
		return nil, false
	}

	current := s.Head

	if current == nil {
		return nil, false
	}

	if position == 0 {
		return current, true
	}

	if position == s.Length-1 {
		return s.Tail, true
	}

	i := 0
	for {
		if i == position {
			return current, true
		}

		current = current.Next

		if current == nil {
			return nil, false
		}

		i++
	}
}

// Set updates the value of the node at the specified position in the list.
// Returns:
// - true if the value was successfully updated
// - false if the index is out of bounds or the list is empty
//
// Time complexity: O(n) in worst case as it uses Get().
func (s *SinglyLinkedList[T]) Set(index int, value T) bool {
	node, ok := s.Get(index)
	if !ok {
		return false
	}

	node.Value = value
	return true
}

// Insert inserts a new node with the given value at the specified index.
// Shifts the existing node at that position (and subsequent ones) to the right.
// Special cases:
// - index == 0 → behaves like Unshift
// - index == Length → behaves like Push
// Returns false if index is out of bounds (<0 or >Length).
//
// Time complexity: O(n) due to traversal when inserting in the middle.
func (s *SinglyLinkedList[T]) Insert(index int, value T) bool {
	if index > s.Length || index < 0 {
		return false
	}

	if index == 0 {
		s.Unshift(value)
		return true
	}

	if index == s.Length {
		s.Push(value)
		return true
	}

	nodeBefore, _ := s.Get(index - 1)

	newNode := &Node[T]{
		Next:  nodeBefore.Next,
		Value: value,
	}

	nodeBefore.Next = newNode
	s.Length++

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
// Special cases:
// - Returns nil if the list is empty
// - Updates Tail if removing the last element
// Time complexity: O(1).
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
// Time complexity: O(n).
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
