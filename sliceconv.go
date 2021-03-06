package sliceconv

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// Atoi converts the provided strings into their respective base 10 integer
// representations. If any of the strings are not valid digits, an error is returned.
func Atoi(str []string) ([]int, error) {
	var ints []int

	for _, s := range str {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.Wrap(err, "one or more strings could not be converted to type int")
		}

		ints = append(ints, i)
	}

	return ints, nil
}

// Itoa converts the provided integers into their respective string representations
// in base 10.
func Itoa(ints []int) []string {
	var str []string

	for _, i := range ints {
		str = append(str, strconv.Itoa(i))
	}
	return str
}

// Atof converts the provided strings into their respective float representations.
// If any of the strings are not valid floats, an error is returned.
func Atof(str []string) ([]float64, error) {
	var flts []float64

	for _, s := range str {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, errors.Wrap(err, "one or more strings could not be converted to type float64")
		}

		flts = append(flts, f)
	}

	return flts, nil
}

// Ftoa converts the provided floats into their respective string representations
// with no exponent, the smallest precision needed, and a bitsize of 64.
func Ftoa(flts []float64) []string {
	var str []string

	for _, f := range flts {
		str = append(str, strconv.FormatFloat(f, 'f', -1, 64))
	}

	return str
}

// Atob converts the provided strings into their respective bool representations.
// Strings that evaluate to true include "1", "t", and "true".
// Strings that evaluate to false include "0", "f", and "false".
// The strings are NOT case-sensitive.
// If any string is invalid, an error is returned.
func Atob(str []string) ([]bool, error) {
	var bools []bool

	for _, s := range str {
		b, err := strconv.ParseBool(strings.ToLower(s))
		if err != nil {
			return nil, errors.Wrap(err, "one or more strings could not be converted to type bool")
		}

		bools = append(bools, b)
	}

	return bools, nil
}

// Btoa converts the provided bools into their respective string representations.
func Btoa(bools []bool) []string {
	var str []string

	for _, b := range bools {
		str = append(str, strconv.FormatBool(b))
	}

	return str
}
