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

// Method to check whether a Stack is empty
func (s *Stack) IsEmpty() bool {
	return s.IsNil() || s.list.Head() == nil
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
