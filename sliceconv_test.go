package sliceconv

import (
	"github.com/pkg/errors"
	"reflect"
	"strconv"
	"testing"
)

func TestAtoi(t *testing.T) {
	tests := []struct {
		name     string
		strings  []string
		wantInts []int
		wantErr  error
	}{
		{"Nil slice", nil, nil, nil},
		{"Empty string slice", []string{}, nil, nil},
		{"Single valid string", []string{"1"}, []int{1}, nil},
		{"Multiple valid strings", []string{"1", "2", "3"}, []int{1, 2, 3}, nil},
		{"Single invalid string", []string{"abc"}, nil, &strconv.NumError{Func: "Atoi", Num: "abc", Err: strconv.ErrSyntax}},
		{"Multiple invalid strings", []string{"abc", "def", "ghi"}, nil, &strconv.NumError{Func: "Atoi", Num: "abc", Err: strconv.ErrSyntax}},
		{"Mixed valid and invalid strings", []string{"123", "abc"}, nil, &strconv.NumError{Func: "Atoi", Num: "abc", Err: strconv.ErrSyntax}},
		{"Out of range", []string{"18446744073709552001"}, nil, &strconv.NumError{Func: "Atoi", Num: "18446744073709552001", Err: strconv.ErrRange}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ints, err := Atoi(test.strings)
			if !reflect.DeepEqual(errors.Cause(err), test.wantErr) {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(ints, test.wantInts) {
				t.Errorf("got: <%v>, want: <%v>", ints, test.wantInts)
			}
		})
	}
}

func TestItoa(t *testing.T) {
	tests := []struct {
		name        string
		ints        []int
		wantStrings []string
	}{
		{"Nil slice", nil, nil},
		{"Empty int slice", []int{}, nil},
		{"Single int", []int{1}, []string{"1"}},
		{"Multiple ints", []int{1, 2, 3}, []string{"1", "2", "3"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := Itoa(test.ints)

			if !reflect.DeepEqual(s, test.wantStrings) {
				t.Errorf("got: <%v>, want: <%v>", s, test.wantStrings)
			}
		})
	}
}
