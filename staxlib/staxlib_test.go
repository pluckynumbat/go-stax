package staxlib

import (
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

func TestIsEmptyNilStack(t *testing.T) {

	var s *Stack

	want := true
	got := s.IsEmpty()

	if got != want {
		t.Errorf("IsEmpty returned incorrected results, want: %v, got: %v", want, got)
	}
}

func TestIsEmptyNonNilStack(t *testing.T) {

	s := &Stack{}

	want := true
	got := s.IsEmpty()

	if got != want {
		t.Errorf("IsEmpty returned incorrected results, want: %v, got: %v", want, got)
	}

	l := listlib.LinkedList{}
	s2 := &Stack{l}

	want = true
	got = s2.IsEmpty()

	if got != want {
		t.Errorf("IsEmpty returned incorrected results, want: %v, got: %v", want, got)
	}
}
