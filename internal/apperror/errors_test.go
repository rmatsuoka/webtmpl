package apperror

import (
	"fmt"
	"io"
	"testing"
)

func Test_lookup(t *testing.T) {
	tests := []struct {
		err  error
		want error
	}{
		{nil, nil},
		{fmt.Errorf("%w", ErrNotFound), ErrNotFound},
		{myError{ErrNotFound}, ErrNotFound},
		{io.EOF, nil},
	}

	for _, test := range tests {
		got, ok := lookup(test.err)
		if test.want == nil {
			if ok {
				t.Errorf("lookup(%v) returns (%v, %t), want (<nil>, false)", test.err, got, ok)
			}
			continue
		}

		if test.want != got {
			t.Errorf("lookup(%v) returns (%v, %t), want (%v, %t)", test.err, got, ok, test.want, test.want != nil)
		}
	}
}

type myError struct {
	err error
}

func (myError) Error() string {
	return "myError"
}

func (e myError) Is(err error) bool {
	return e.err == err
}
