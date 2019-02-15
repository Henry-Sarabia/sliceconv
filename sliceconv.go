package sliceconv

import (
	"github.com/pkg/errors"
	"strconv"
)

func Itoa(ints ...int) []string {
	var str []string
	for _, i := range ints {
		str = append(str, strconv.Itoa(i))
	}
	return str
}

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
