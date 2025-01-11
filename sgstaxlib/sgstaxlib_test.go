package sgstaxlib

import (
	"fmt"
	"testing"

	"github.com/pluckynumbat/linked-list-stuff-go/sdlistlib"
)

type prInt int

func (p *prInt) String() string {
	return fmt.Sprintf("%v", *p)
}

func TestIsNil(t *testing.T) {

	var s1 *SemiGenericStack[*prInt]
	s2 := &SemiGenericStack[*prInt]{}

	var tests = []struct {
		name  string
		stack *SemiGenericStack[*prInt]
		want  bool
	}{
		{"nil true", s1, true},
		{"nil false", s2, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.stack.IsNil()
			want := test.want

			if got != want {
				t.Errorf("IsNil() gave incorrect results, want: %v, got : %v", want, got)
			}
		})
	}
}

func TestIsListNil(t *testing.T) {

	var s1 *SemiGenericStack[*prInt]
	s2 := &SemiGenericStack[*prInt]{}

	l := &sdlistlib.SemiGenericList[*prInt]{}
	s3 := &SemiGenericStack[*prInt]{l}

	var tests = []struct {
		name  string
		stack *SemiGenericStack[*prInt]
		want  bool
	}{
		{"nil stack", s1, true},
		{"nil list", s2, true},
		{"nil false", s3, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.stack.isListNil()
			want := test.want

			if got != want {
				t.Errorf("isListNil() gave incorrect results, want: %v, got : %v", want, got)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {

	var s1 *SemiGenericStack[*prInt]
	s2 := &SemiGenericStack[*prInt]{}

	l := &sdlistlib.SemiGenericList[*prInt]{}
	s3 := &SemiGenericStack[*prInt]{l}

	var pr prInt = 1
	var ptr = &pr
	l2 := &sdlistlib.SemiGenericList[*prInt]{}
	l2.AddAtBeginning(ptr)
	s4 := &SemiGenericStack[*prInt]{l2}

	var tests = []struct {
		name  string
		stack *SemiGenericStack[*prInt]
		want  bool
	}{
		{"nil stack", s1, true},
		{"non nil stack, nil list", s2, true},
		{"non nil stack, empty list", s3, true},
		{"non nil stack, non empty list", s4, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.stack.IsEmpty()
			want := test.want

			if got != want {
				t.Errorf("IsEmpty() gave incorrect results, want: %v, got : %v", want, got)
			}
		})
	}
}

func TestPeek(t *testing.T) {
	var s1 *SemiGenericStack[*prInt]
	s2 := &SemiGenericStack[*prInt]{}

	l := &sdlistlib.SemiGenericList[*prInt]{}
	s3 := &SemiGenericStack[*prInt]{l}

	var pr prInt = 1
	var ptr = &pr
	l2 := &sdlistlib.SemiGenericList[*prInt]{}
	l2.AddAtBeginning(ptr)
	s4 := &SemiGenericStack[*prInt]{l2}

	var tests = []struct {
		name         string
		stack        *SemiGenericStack[*prInt]
		expValString string
		expErr       error
	}{
		{"nil stack", s1, "nil", stackNilError},
		{"non-nil stack, nil list", s2, "nil", stackEmptyError},
		{"empty stack", s3, "nil", stackEmptyError},
		{"non-empty stack", s4, "1", nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			gotVal, gotErr := test.stack.Peek()
			wantErr := test.expErr

			gotString := gotVal.String()
			wantString := test.expValString

			if gotErr != wantErr {
				t.Errorf("Unexpected error for Peek(), want: %v, got : %v", wantErr, gotErr)
			}

			if gotString != wantString {
				t.Errorf("Incorrect result for Peek(), want: %v, got : %v", wantString, gotString)
			}
		})
	}
}

func TestPeekStackTillEmpty(t *testing.T) {
	l := &sdlistlib.SemiGenericList[*prInt]{}

	var pr1, pr2, pr3 prInt = 1, 2, 3
	ptrs := []*prInt{&pr1, &pr2, &pr3}

	for _, v := range ptrs {
		err := l.AddAtBeginning(v)
		if err != nil {
			t.Fatalf("list's AddAtBeginning() failed with error: %v", err)
		}
	}

	s := &SemiGenericStack[*prInt]{l}

	var tests = []struct {
		name          string
		wantValString string
	}{
		{"3 element stack", "3"},
		{"2 element stack", "2"},
		{"1 element stack", "1"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			val, err := s.Peek()
			if err != nil {
				t.Fatalf("Peek() failed with error: %v", err)
			}

			got := val.String()
			want := test.wantValString

			if got != want {
				t.Errorf("Incorrect result for Peek(), want: %v, got : %v", want, got)
			}

			_, err2 := l.RemoveFirst()
			if err != nil {
				t.Fatalf("list's RemoveFirst() failed with error: %v", err2)
			}
		})
	}

	// check is empty
	isEmpty := s.IsEmpty()
	if !isEmpty {
		t.Errorf("stack should be empty after removing all elements from the underlying list")
	}
}
