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

