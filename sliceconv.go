package sliceconv

import (
	"github.com/pkg/errors"
	"strconv"
)

// Atoi converts the provided strings into their respective base 10 integer
// representations. If any of the strings are not valid digits, an error is returned.
func Atoi(str ...string) ([]int, error) {
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
func Itoa(ints ...int) []string {
	var str []string
	for _, i := range ints {
		str = append(str, strconv.Itoa(i))
	}
	return str
}
