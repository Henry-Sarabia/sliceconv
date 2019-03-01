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
		name    string
		ints    []int
		wantStr []string
	}{
		{"Nil slice", nil, nil},
		{"Empty int slice", []int{}, nil},
		{"Single int", []int{1}, []string{"1"}},
		{"Multiple ints", []int{1, 2, 3}, []string{"1", "2", "3"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := Itoa(test.ints)

			if !reflect.DeepEqual(s, test.wantStr) {
				t.Errorf("got: <%v>, want: <%v>", s, test.wantStr)
			}
		})
	}
}

func TestAtof(t *testing.T) {
	tests := []struct {
		name     string
		strings  []string
		wantFlts []float64
		wantErr  error
	}{
		{"Nil slice", nil, nil, nil},
		{"Empty string slice", []string{}, nil, nil},
		{"Single valid string", []string{"0.000000000000001"}, []float64{0.000000000000001}, nil},
		{"Multiple valid strings", []string{"1.25", "2.125", "3.0625", "4.03125", "5.015625"}, []float64{1.25, 2.125, 3.0625, 4.03125, 5.015625}, nil},
		{"Single invalid string", []string{"abc"}, nil, &strconv.NumError{Func: "ParseFloat", Num: "abc", Err: strconv.ErrSyntax}},
		{"Multiple invalid strings", []string{"abc", "def", "ghi"}, nil, &strconv.NumError{Func: "ParseFloat", Num: "abc", Err: strconv.ErrSyntax}},
		{"Mixed valid and invalid strings", []string{"1.25", "abc"}, nil, &strconv.NumError{Func: "ParseFloat", Num: "abc", Err: strconv.ErrSyntax}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			flts, err := Atof(test.strings)
			if !reflect.DeepEqual(errors.Cause(err), test.wantErr) {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(flts, test.wantFlts) {
				t.Errorf("got: <%v>, want: <%v>", flts, test.wantFlts)
			}
		})
	}
}

func TestFtoa(t *testing.T) {
	tests := []struct {
		name    string
		flts    []float64
		wantStr []string
	}{
		{"Nil slice", nil, nil},
		{"Empty float slice", []float64{}, nil},
		{"Single float", []float64{1.5}, []string{"1.5"}},
		{"Multiple floats", []float64{1.25, 2.125, 3.0625, 4.03125, 5.015625}, []string{"1.25", "2.125", "3.0625", "4.03125", "5.015625"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := Ftoa(test.flts)

			if !reflect.DeepEqual(s, test.wantStr) {
				t.Errorf("got: <%v>, want: <%v>", s, test.wantStr)
			}
		})
	}
}

func TestAtob(t *testing.T) {
	tests := []struct {
		name      string
		strings   []string
		wantBools []bool
		wantErr   error
	}{
		{"String '1'", []string{"1"}, []bool{true}, nil},
		{"String 't", []string{"t"}, []bool{true}, nil},
		{"String 'T'", []string{"T"}, []bool{true}, nil},
		{"String 'TRUE'", []string{"TRUE"}, []bool{true}, nil},
		{"String 'true'", []string{"true"}, []bool{true}, nil},
		{"String 'True'", []string{"True"}, []bool{true}, nil},
		{"String '0'", []string{"0"}, []bool{false}, nil},
		{"String 'f'", []string{"f"}, []bool{false}, nil},
		{"String 'F'", []string{"F"}, []bool{false}, nil},
		{"String 'FALSE'", []string{"FALSE"}, []bool{false}, nil},
		{"String 'false'", []string{"false"}, []bool{false}, nil},
		{"String 'False'", []string{"False"}, []bool{false}, nil},
		{"Nil slice", nil, nil, nil},
		{"Empty string slice", []string{}, nil, nil},
		{"Single valid string", []string{"true"}, []bool{true}, nil},
		{"Multiple valid strings", []string{"1", "0", "t", "f"}, []bool{true, false, true, false}, nil},
		{"Single invalid string", []string{"invalid"}, nil, &strconv.NumError{Func: "ParseBool", Num: "invalid", Err: strconv.ErrSyntax}},
		{"Multiple invalid strings", []string{"invalid", "trueeee", "faaaalse"}, nil, &strconv.NumError{Func: "ParseBool", Num: "invalid", Err: strconv.ErrSyntax}},
		{"Mixed valid and invalid strings", []string{"true", "invalid", "false"}, nil, &strconv.NumError{Func: "ParseBool", Num: "invalid", Err: strconv.ErrSyntax}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bools, err := Atob(test.strings)
			if !reflect.DeepEqual(errors.Cause(err), test.wantErr) {
				t.Errorf("got: <%v>, want: <%v>", errors.Cause(err), test.wantErr)
			}

			if !reflect.DeepEqual(bools, test.wantBools) {
				t.Errorf("got: <%v>, want: <%v>", bools, test.wantBools)
			}
		})
	}
}

func TestBtoa(t *testing.T) {
	tests := []struct {
		name    string
		bools   []bool
		wantStr []string
	}{
		{"Nil slice", nil, nil},
		{"Empty bool slice", []bool{}, nil},
		{"Single true", []bool{true}, []string{"true"}},
		{"Multiple true", []bool{true, true, true}, []string{"true", "true", "true"}},
		{"Single false", []bool{false}, []string{"false"}},
		{"Multiple false", []bool{false, false, false}, []string{"false", "false", "false"}},
		{"Mixed true and false", []bool{true, false, true, false}, []string{"true", "false", "true", "false"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := Btoa(test.bools)

			if !reflect.DeepEqual(s, test.wantStr) {
				t.Errorf("got: <%v>, want: <%v>", s, test.wantStr)
			}
		})
	}
}
