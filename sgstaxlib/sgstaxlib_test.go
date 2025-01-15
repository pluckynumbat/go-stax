package sgstaxlib

import (
	"fmt"
	"testing"

	"github.com/pluckynumbat/linked-list-stuff-go/sdlistlib"
)

type prInt int       // printable int
type prString string // printable string

func (p prInt) String() string {
	return fmt.Sprintf("%v", int(p))
}

func (p prString) String() string {
	return fmt.Sprintf("%v", string(p))
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

			if gotErr != wantErr {
				t.Errorf("Unexpected error for Peek(), want: %v, got : %v", wantErr, gotErr)
			} else if gotErr != nil {
				fmt.Println(gotErr)
			} else {
				gotString := gotVal.String()
				wantString := test.expValString
				if gotString != wantString {
					t.Errorf("Incorrect result for Peek(), want: %v, got : %v", wantString, gotString)
				}
			}
		})
	}

	// peek stack till empty
	l3 := &sdlistlib.SemiGenericList[*prInt]{}

	var pr1, pr2, pr3 prInt = 1, 2, 3
	ptrs := []*prInt{&pr1, &pr2, &pr3}

	for _, v := range ptrs {
		err := l3.AddAtBeginning(v)
		if err != nil {
			t.Fatalf("list's AddAtBeginning() failed with error: %v", err)
		}
	}

	s := &SemiGenericStack[*prInt]{l3}

	var tests2 = []struct {
		name          string
		wantValString string
	}{
		{"3 element stack", "3"},
		{"2 element stack", "2"},
		{"1 element stack", "1"},
	}

	for _, test := range tests2 {
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

			_, err2 := l3.RemoveFirst()
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

func TestPush(t *testing.T) {

	// Part 1: Stack of prInt pointers
	var s1 *SemiGenericStack[*prInt]

	var pr1, pr2, pr3 prInt = 1, 2, 3

	err := s1.Push(&pr1)
	if err == nil {
		t.Error("Calling Push() on a nil stack should return an error!")
	} else {
		fmt.Println(err)
	}

	s1 = &SemiGenericStack[*prInt]{}

	var tests1 = []struct {
		name   string
		val    *prInt
		expVal string
	}{
		{"push to empty prInt pointer stack", new(prInt), "0"},
		{"push to 1 element prInt pointer stack", &pr1, "1"},
		{"push to 2 element prInt pointer stack", &pr2, "2"},
		{"push to 3 element prInt pointer stack", &pr3, "3"},
	}

	for _, test := range tests1 {
		t.Run(test.name, func(t *testing.T) {
			err := s1.Push(test.val)
			if err != nil {
				t.Errorf("Push() failed with error: %v", err)
			} else {
				peekVal, err2 := s1.Peek()
				if err2 != nil {
					t.Errorf("Peek() failed with error: %v", err2)
				}

				got := peekVal.String()
				want := test.expVal

				if got != want {
					t.Errorf("Push() returned incorrect results, want: %v, got: %v", want, got)
				}
			}
		})
	}

	// Part 2: Stack of prString pointers
	var s2 *SemiGenericStack[*prString]

	var prS0 prString
	var prS1, prS2, prS3 prString = "a", "b", "c"

	err = s2.Push(&prS1)
	if err == nil {
		t.Error("Calling Push() on a nil stack should return an error!")
	} else {
		fmt.Println(err)
	}

	s2 = &SemiGenericStack[*prString]{}

	var tests2 = []struct {
		name   string
		val    *prString
		expVal string
	}{
		{"push to empty prString stack", &prS0, ""},
		{"push to 1 element prString stack", &prS1, "a"},
		{"push to 2 element prString stack", &prS2, "b"},
		{"push to 3 element prString stack", &prS3, "c"},
	}

	for _, test := range tests2 {
		t.Run(test.name, func(t *testing.T) {
			err := s2.Push(test.val)
			if err != nil {
				t.Errorf("Push() failed with error: %v", err)
			} else {
				peekVal, err2 := s2.Peek()
				if err2 != nil {
					t.Errorf("Peek() failed with error: %v", err2)
				}

				got := peekVal.String()
				want := test.expVal

				if got != want {
					t.Errorf("Push() returned incorrect results, want: %v, got: %v", want, got)
				}
			}
		})
	}
}

func TestPopNilStack(t *testing.T) {
	var s1 *SemiGenericStack[prInt]

	_, err := s1.Pop()
	if err == nil {
		t.Errorf("Pop() on a nil stack should return an error")
	} else {
		fmt.Println(err)
	}

	var s2 *SemiGenericStack[*prString]

	_, err = s2.Pop()
	if err == nil {
		t.Errorf("Pop() on a nil stack should return an error")
	} else {
		fmt.Println(err)
	}
}

func TestPopEmptyStack(t *testing.T) {
	s1 := &SemiGenericStack[*prInt]{}

	_, err := s1.Pop()
	if err == nil {
		t.Errorf("Pop() on an empty stack should return an error")
	} else {
		fmt.Println(err)
	}

	s2 := &SemiGenericStack[prString]{}

	_, err = s2.Pop()
	if err == nil {
		t.Errorf("Pop() on an empty stack should return an error")
	} else {
		fmt.Println(err)
	}
}

func TestPopNilEmptyStacks(t *testing.T) {

	// Stack of prInt
	t.Run("stacks of prInt", func(t *testing.T) {
		var s1 *SemiGenericStack[prInt]
		s2 := &SemiGenericStack[prInt]{}

		l := &sdlistlib.SemiGenericList[prInt]{}
		s3 := &SemiGenericStack[prInt]{l}

		var pr prInt = 1
		l2 := &sdlistlib.SemiGenericList[prInt]{}
		l2.AddAtBeginning(pr)
		s4 := &SemiGenericStack[prInt]{l2}

		var tests = []struct {
			name   string
			stack  *SemiGenericStack[prInt]
			expVal prInt
			expErr error
		}{
			{"nil stack", s1, 0, stackNilError},
			{"non-nil stack, nil list", s2, 0, stackEmptyError},
			{"empty stack", s3, 0, stackEmptyError},
			{"non-empty stack", s4, 1, nil},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {

				val, gotErr := test.stack.Pop()
				if gotErr != test.expErr {
					t.Errorf("Unexpected error for Pop(), want: %v, got : %v", test.expErr, gotErr)
				} else if gotErr != nil {
					fmt.Println(gotErr)
				} else {
					if val != test.expVal {
						t.Errorf("Incorrect result for Pop(), want: %v, got : %v", test.expVal, val)
					}
				}
			})
		}
	})

	// Stack of pointers to prString
	t.Run("stacks of *prString", func(t *testing.T) {
		var s1 *SemiGenericStack[*prString]
		s2 := &SemiGenericStack[*prString]{}

		l := &sdlistlib.SemiGenericList[*prString]{}
		s3 := &SemiGenericStack[*prString]{l}

		var pr prString = "a"
		l2 := &sdlistlib.SemiGenericList[*prString]{}
		l2.AddAtBeginning(&pr)
		s4 := &SemiGenericStack[*prString]{l2}

		var tests = []struct {
			name   string
			stack  *SemiGenericStack[*prString]
			expVal *prString
			expErr error
		}{
			{"nil stack", s1, nil, stackNilError},
			{"non-nil stack, nil list", s2, nil, stackEmptyError},
			{"empty stack", s3, nil, stackEmptyError},
			{"non-empty stack", s4, &pr, nil},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {

				val, gotErr := test.stack.Pop()
				if gotErr != test.expErr {
					t.Errorf("Unexpected error for Pop(), want: %v, got : %v", test.expErr, gotErr)
				} else if gotErr != nil {
					fmt.Println(gotErr)
				} else {
					if val != test.expVal {
						t.Errorf("Incorrect result for Pop(), want: %v, got : %v", test.expVal, val)
					}
				}
			})
		}
	})
}

func TestPopTillEmpty(t *testing.T) {

	//Stack of pointers to prInt
	t.Run("stack of *prInt", func(t *testing.T) {
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
			name       string
			wantVal    *prInt
			newTop     *prInt
			expPeekErr error
		}{
			{"3 element stack", &pr3, &pr2, nil},
			{"2 element stack", &pr2, &pr1, nil},
			{"1 element stack", &pr1, nil, stackEmptyError},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {

				val, err := s.Pop()
				if err != nil {
					t.Errorf("Pop() failed with error: %v", err)
				} else {
					if val != test.wantVal {
						t.Errorf("Incorrect result for Pop(), want: %v, got : %v", test.wantVal, val)
					} else {
						topVal, peekErr := s.Peek()
						if peekErr != test.expPeekErr {
							t.Errorf("Unexpected error for Peek(), want: %v, got : %v", test.expPeekErr, peekErr)
						} else if peekErr != nil {
							fmt.Println(peekErr)
						} else if topVal != test.newTop {
							t.Errorf("Incorrect result for Peek(), want: %v, got : %v", test.newTop, topVal)
						}
					}
				}
			})
		}
	})

	//Stack of prString
	t.Run("stack of prString", func(t *testing.T) {
		l := &sdlistlib.SemiGenericList[prString]{}

		prStrs := []prString{"a", "b", "c"}

		for _, v := range prStrs {
			err := l.AddAtBeginning(v)
			if err != nil {
				t.Fatalf("list's AddAtBeginning() failed with error: %v", err)
			}
		}

		s := &SemiGenericStack[prString]{l}

		var tests = []struct {
			name       string
			wantVal    prString
			newTop     prString
			expPeekErr error
		}{
			{"3 element stack", "c", "b", nil},
			{"2 element stack", "b", "a", nil},
			{"1 element stack", "a", "", stackEmptyError},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {

				val, err := s.Pop()
				if err != nil {
					t.Errorf("Pop() failed with error: %v", err)
				} else {
					if val != test.wantVal {
						t.Errorf("Incorrect result for Pop(), want: %v, got : %v", test.wantVal, val)
					} else {
						topVal, peekErr := s.Peek()
						if peekErr != test.expPeekErr {
							t.Errorf("Unexpected error for Peek(), want: %v, got : %v", test.expPeekErr, peekErr)
						} else if peekErr != nil {
							fmt.Println(peekErr)
						} else if topVal != test.newTop {
							t.Errorf("Incorrect result for Peek(), want: %v, got : %v", test.newTop, topVal)
						}
					}
				}
			})
		}
	})
}

