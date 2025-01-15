// Semi generic stack that contain data of any type that implements the fmt.Stringer interface
package sgstaxlib

import (
	"fmt"

	"github.com/pluckynumbat/linked-list-stuff-go/sdlistlib"
)

var stackNilError = fmt.Errorf("The stack is nil")
var stackEmptyError = fmt.Errorf("The stack is empty")

type SemiGenericStack[T fmt.Stringer] struct {
	sdlist *sdlistlib.SemiGenericList[T]
}

// Method to check whether a pointer to a Semi Generic Stack is nil
func (stack *SemiGenericStack[T]) IsNil() bool {
	return stack == nil
}

// Internal Method to check whether the underlying list is nil
func (stack *SemiGenericStack[T]) isListNil() bool {
	return stack.IsNil() || stack.sdlist.IsNil()
}

// Method to check whether a Semi Generic Stack is empty
func (stack *SemiGenericStack[T]) IsEmpty() bool {
	return stack.IsNil() || stack.isListNil() || stack.sdlist.Head() == nil
}

// Method to check the the data at the top of the Semi Generic Stack
func (stack *SemiGenericStack[T]) Peek() (T, error) {

	if stack.IsNil() {
		return *new(T), stackNilError
	}

	if stack.IsEmpty() {
		return *new(T), stackEmptyError
	}

	data, err := stack.sdlist.Head().GetData()
	if err != nil {
		return *new(T), fmt.Errorf("Error retrieving data from top:  %v", err)
	}

	return data, nil
}

// Method to add a new element to the top of the Semi Generic Stack
func (stack *SemiGenericStack[T]) Push(val T) error {
	if stack.IsNil() {
		return stackNilError
	}

	if stack.isListNil() {
		stack.sdlist = &sdlistlib.SemiGenericList[T]{}
	}

	stack.sdlist.AddAtBeginning(val)

	return nil
}

// Method to remove the top most element from a Semi Generic Stack (if present) and return it
func (stack *SemiGenericStack[T]) Pop() (T, error) {
	if stack.IsNil() {
		return *new(T), stackNilError
	}

	if stack.IsEmpty() {
		return *new(T), stackEmptyError
	}

	node, err := stack.sdlist.RemoveFirst()
	if err != nil {
		return *new(T), fmt.Errorf("Pop() failed with error: %v", err)
	}

	val, err := node.GetData()
	if err != nil {
		return *new(T), fmt.Errorf("Pop() failed with error: %v", err)
	}

	return val, nil
}
