package staxlib

import (
	"fmt"
	"testing"

	"github.com/pluckynumbat/linked-list-stuff-go/listlib"
)

func TestIsNil(t *testing.T) {
	var s1, s2 *Stack
	s2 = &Stack{}
	var tests = []struct {
		name string
		s    *Stack
		want bool
	}{
		{"nil true", s1, true},
		{"nil false", s2, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want := test.want
			got := test.s.IsNil()

			if want != got {
				t.Errorf("IsNil returned incorrected results, want: %v, got: %v", want, got)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	var s1 *Stack

	s2 := &Stack{}

	emptyList := &listlib.LinkedList{}
	s3 := &Stack{emptyList}

	nonEmptyList := &listlib.LinkedList{}
	nonEmptyList.AddToBeginning("a")
	s4 := &Stack{nonEmptyList}

	var tests = []struct {
		name string
		s    *Stack
		want bool
	}{
		{"nil stack", s1, true},
		{"non nil stack, nil list", s2, true},
		{"non nil stack, empty list", s3, true},
		{"non nil stack, non empty list", s4, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want := test.want
			got := test.s.IsEmpty()

			if want != got {
				t.Errorf("IsEmpty returned incorrected results, want: %v, got: %v", want, got)
			}
		})
	}
}

func TestPeekNilStack(t *testing.T) {
	var s *Stack
	_, err := s.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		t.Errorf("Calling Peek() on a nil stack should return an error!")
	}
}

func TestPeekEmptyStack(t *testing.T) {
	s := &Stack{}
	_, err := s.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		t.Errorf("Calling Peek() on an empty stack should return an error!")
	}
}

