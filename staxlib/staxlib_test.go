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

func TestIsListNil(t *testing.T) {
	var s1, s2, s3 *Stack
	s2 = &Stack{}

	l := &listlib.LinkedList{}
	s3 = &Stack{l}
	var tests = []struct {
		name string
		s    *Stack
		want bool
	}{
		{"nil stack", s1, true},
		{"nil list", s2, true},
		{"nil false", s3, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want := test.want
			got := test.s.isListNil()

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

func TestPeekStackEmptyList(t *testing.T) {
	emptyList := &listlib.LinkedList{}
	s := &Stack{emptyList}
	_, err := s.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		t.Errorf("Calling Peek() on an empty stack should return an error!")
	}
}

func TestPeekStackNonEmptyList(t *testing.T) {
	nonEmptyList := &listlib.LinkedList{}
	nonEmptyList.AddToBeginning("a")
	s := &Stack{nonEmptyList}
	data, err := s.Peek()
	if err != nil {
		t.Errorf("Peek() on the Stack failed, error: %v", err)
	} else {
		want := "a"
		got := data
		if got != want {
			t.Errorf("Incorrect results for Peek() on the Stack, want: %v, got: %v", want, got)
		}
	}
}

func TestPeekStackTillEmpty(t *testing.T) {
	l := listlib.ConstructFromValues("a", "b", "c")
	lp := &l
	s := &Stack{lp}

	var tests = []struct {
		name string
		want string
	}{
		{"3 elements", "a"},
		{"2 elements", "b"},
		{"1 element", "c"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, err := s.Peek()
			if err != nil {
				t.Errorf("Peek() on the Stack failed, error: %v", err)
			} else {
				want := test.want
				got := data
				if got != want {
					t.Errorf("Incorrect results for Peek() on the Stack, want: %v, got: %v", want, got)
				}
			}
			l.RemoveAtBeginning()
		})
	}

	_, err := s.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		t.Errorf("Calling Peek() on an empty stack should return an error!")
	}
}

func TestPush(t *testing.T) {
	var s0 *Stack
	err := s0.Push("a")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Errorf("Calling Push() on a nil stack should return an error!")
	}

	s := &Stack{}
	var tests = []struct {
		name    string
		pushVal string
		want    string
	}{
		{"1 element", "a", "a"},
		{"2 elements", "b", "b"},
		{"3 elements", "c", "c"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err = s.Push(test.pushVal)
			if err != nil {
				t.Errorf("Push() on the Stack failed, error: %v", err)
			} else {
				data, err2 := s.Peek()
				if err2 != nil {
					t.Errorf("Peek() on the Stack failed, error: %v", err2)
				} else {
					want := test.want
					got := data
					if got != want {
						t.Errorf("Incorrect results for Push() on the Stack, want: %v, got: %v", want, got)
					}
				}
			}
		})
	}
}

func TestPopNilStack(t *testing.T) {
	var s *Stack
	_, err := s.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		t.Errorf("Calling Pop() on a nil stack should return an error!")
	}
}

func TestPopEmptyStack(t *testing.T) {
	s := &Stack{}
	_, err := s.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		t.Errorf("Calling Pop() on an empty stack should return an error!")
	}
}

func TestPopStackTillEmpty(t *testing.T) {
	l := listlib.ConstructFromValues("a", "b", "c")
	s := &Stack{&l}

	var tests = []struct {
		name       string
		popVal     string
		newTop     string
		expPeekErr error
	}{
		{"3 elements", "a", "b", nil},
		{"2 elements", "b", "c", nil},
		{"1 element", "c", "", stackEmptyError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val, err := s.Pop()
			if err != nil {
				t.Errorf("Pop() on the Stack failed, error: %v", err)
			} else {
				want := test.popVal
				got := val

				if got != want {
					t.Errorf("Incorrect results for Pop() on the Stack, want: %v, got: %v", want, got)
				}

				val2, err2 := s.Peek()

				if test.expPeekErr == nil {
					if err2 != nil {
						t.Errorf("Peek() on the Stack failed unexpectedly, error: %v", err2)
					}
				} else {
					if err2 != test.expPeekErr {
						t.Errorf("Peek() error: %v doesn't match expected error: %v", err2, test.expPeekErr)
					} else {
						fmt.Println(err2)
					}
				}

				got = val2
				want = test.newTop
				if got != want {
					t.Errorf("Incorrect results for Peek() on the Stack, want: %v, got: %v", want, got)
				}
			}
		})
	}
}

func TestStackOperations(t *testing.T) {
	s := &Stack{}
	var pushTests = []struct {
		name    string
		pushVal string
		want    string
	}{
		{"push a", "a", "a"},
		{"push b", "b", "b"},
		{"push c", "c", "c"},
	}

	for _, test := range pushTests {
		t.Run(test.name, func(t *testing.T) {
			err := s.Push(test.pushVal)
			if err != nil {
				t.Errorf("Push() on the Stack failed, error: %v", err)
			} else {
				data, err2 := s.Peek()
				if err2 != nil {
					t.Errorf("Peek() on the Stack failed, error: %v", err2)
				} else {
					want := test.want
					got := data
					if got != want {
						t.Errorf("Incorrect results for Push() on the Stack, want: %v, got: %v", want, got)
					}
				}
			}
		})
	}

	var popTests = []struct {
		name       string
		popVal     string
		newTop     string
		expPeekErr error
	}{
		{"pop c", "c", "b", nil},
		{"pop b", "b", "a", nil},
		{"pop a", "a", "", stackEmptyError},
	}

	for _, test := range popTests {
		t.Run(test.name, func(t *testing.T) {
			val, err := s.Pop()
			if err != nil {
				t.Errorf("Pop() on the Stack failed, error: %v", err)
			} else {
				want := test.popVal
				got := val

				if got != want {
					t.Errorf("Incorrect results for Pop() on the Stack, want: %v, got: %v", want, got)
				}

				val2, err2 := s.Peek()

				if test.expPeekErr == nil {
					if err2 != nil {
						t.Errorf("Peek() on the Stack failed unexpectedly, error: %v", err2)
					}
				} else {
					if err2 != test.expPeekErr {
						t.Errorf("Peek() error: %v doesn't match expected error: %v", err2, test.expPeekErr)
					} else {
						fmt.Println(err2)
					}
				}

				got = val2
				want = test.newTop
				if got != want {
					t.Errorf("Incorrect results for Peek() on the Stack, want: %v, got: %v", want, got)
				}
			}
		})
	}

	var stateTests = []struct {
		name   string
		fnName string
		fn     func() bool
		want   bool
	}{
		{"is nil", "IsNil()", s.IsNil, false},
		{"is list nil", "isListNil()", s.isListNil, false},
		{"is empty", "IsEmpty()", s.IsEmpty, true},
	}

	for _, test := range stateTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.fn()
			want := test.want
			if got != want {
				t.Errorf("Incorrect results for state check function %v on the Stack, want: %v, got: %v", test.fnName, want, got)
			}
		})
	}

}
