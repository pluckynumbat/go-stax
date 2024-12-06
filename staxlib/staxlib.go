package staxlib

import (
	"github.com/pluckynumbat/linked-list-stuff-go/listlib"
)

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
