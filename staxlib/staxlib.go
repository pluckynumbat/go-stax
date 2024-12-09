package staxlib

import (
	"fmt"

	"github.com/pluckynumbat/linked-list-stuff-go/listlib"
)

var stackNilError = fmt.Errorf("The stack is nil")
var stackEmptyError = fmt.Errorf("The stack is empty")

type Stack struct {
	list *listlib.LinkedList
}

// Method to check whether a pointer to a Stack is nil
func (s *Stack) IsNil() bool {
	return s == nil
}

// Internal Method to check whether the underlying list is nil
func (s *Stack) isListNil() bool {
	return s.IsNil() || s.list.IsNil()
}

// Method to check whether a Stack is empty
func (s *Stack) IsEmpty() bool {
	return s.IsNil() || s.isListNil() || s.list.Head() == nil
}

// Method to check the the data at the top of the Stack
func (s *Stack) Peek() (string, error) {
	if s.IsNil() {
		return "", stackNilError
	}

	if s.IsEmpty() {
		return "", stackEmptyError
	}

	data, err := s.list.Head().GetData()
	if err != nil {
		return "", fmt.Errorf("Error rerieving data from top:  %v", err)
	}

	return data, nil
}

// Method to add a new entry to the top of the Stack
func (s *Stack) Push(value string) error {
	if s.IsNil() {
		return stackNilError
	}

	if s.isListNil() {
		newList := listlib.ConstructFromValues(value)
		s.list = &newList
		return nil
	}

	s.list.AddToBeginning(value)
	return nil
}

// Method to remove an entry from the top of the Stack
func (s *Stack) Pop() (string, error) {
	if s.IsNil() {
		return "", stackNilError
	}

	if s.IsEmpty() {
		return "", stackEmptyError
	}

	node := s.list.RemoveAtBeginning()
	if node == nil {
		return "", fmt.Errorf("Nil pointer encountered while popping from the Stack")
	}

	val, err := node.GetData()
	if err != nil {
		return "", fmt.Errorf("Error encountered while popping from the Stack:  %v", err)
	}
	return val, nil
}
