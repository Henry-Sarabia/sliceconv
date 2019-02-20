# sliceconv

[![GoDoc](https://godoc.org/github.com/Henry-Sarabia/sliceconv?status.svg)](https://godoc.org/github.com/Henry-Sarabia/sliceconv) [![Build Status](https://travis-ci.com/Henry-Sarabia/sliceconv.svg?branch=master)](https://travis-ci.com/Henry-Sarabia/sliceconv) [![Coverage Status](https://coveralls.io/repos/github/Henry-Sarabia/sliceconv/badge.svg?branch=master)](https://coveralls.io/github/Henry-Sarabia/sliceconv?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/Henry-Sarabia/sliceconv)](https://goreportcard.com/report/github.com/Henry-Sarabia/sliceconv)

Sliceconv implements conversions to and from string representations of primitive types on entire slices.
This package takes its cue from the Golang standard library's [strconv](https://golang.org/pkg/strconv/). 

## Installation

If you do not have Go installed yet, you can find installation instructions 
[here](https://golang.org/doc/install).

To pull the most recent version of sliceconv, use `go get`.

```
go get github.com/Henry-Sarabia/sliceconv
```

Then import the package into your project.

```go
import "github.com/Henry-Sarabia/sliceconv"
```

## Usage

### Strings to Integers

As an example, assume you have a list of test scores as a slice of strings. You mean to find the 
highest score by passing the list into a `findMax` function. Unfortunately, the function has the signature 
`findMax([]int) int`. Using sliceconv, simply convert the slice of strings into a slice of ints to
resolve the issue.

```go
scores := []string{"98", "85", "100", "76", "92"}

ints, err := sliceconv.Atoi(scores)
if err != nil {
	// handle error
}

max := findMax(ints)
// output: 100
```

Be aware that the `sliceconv.Atoi` function fulfills the same contract as its counterpart 
`strconv.Atoi`. That is to say, the `Atoi` functions treat the strings as integers in base 10 and
will return an error if any of the resulting integers cannot fit into the regular `int` type.

### Integers to Strings

This time around, assume you still have a list of test scores but as a slice of ints. You mean to
format a class report card by passing the list of scores into a `formatReport` function. With a 
streak of bad luck, the function has the signature `formatReport([]string) string`. With sliceconv,
you can just convert the slice of ints into a slice of strings to satisfy the requirements.

```go
scores := []int{98, 85, 100, 76, 92}

str := sliceconv.Itoa(scores)
report := formatReport(str)
// output: student_1: 98, student_2: 85, ..., student_5: 92
```

Similarly to `sliceconv.Atoi`, the `sliceconv.Itoa` function assumes the integers are in base 10.

### Strings to Floats

Sticking to the test score example, assume you have a list of test scores as a slice of strings.
You mean to find the average score by passing the list into an averaging function with the signature
`findAvg([]float64) float64`. Taking advantage of sliceconv, you can convert the entire list of strings
into a slice of floats.

```go
scores := []string{"98.5", "85.1", "100", "76.9", "92.3"}

flts, err := sliceconv.Atof(scores)
if err != nil {
	// handle error
}

avg := findAvg(flts)
// output: 90.56
```

Be aware that the `sliceconv.Atof` function assumes the stringified floats have a bitsize of 64.

### Floats to Strings

Again, assume you have a list of test scores but as a slice of float64s. You mean to format a class
report card by passing the list of scores into a `formatReport` function. The function has the following
signature: `formatReport([]string) string`. Using sliceconv, simply convert the slice of float64s into
a slice of strings for the `formatReport` function.

```go
scores := []float64{98.5, 85.1, 100, 76.9, 92.3}

str := sliceconv.Ftoa(scores)
report := formatReport(str)
// output: student_1: 98.5, student_2: 85.1, ..., student_5: 92.3
```

Please take note that the `sliceconv.Ftoa` function will fulfill the following contract: the string
representation of the provided floats will have no exponent, the smallest precision needed to properly
represent the number, and a bitsize of 64.

## Contributions

If you would like to contribute to this project, please adhere to the following
guidelines.

* Submit an issue describing the problem.
* Fork the repo and add your contribution.
* Add appropriate tests.
* Run go fmt, go vet, and golint.
* Prefer idiomatic Go over non-idiomatic code.
* Follow the basic Go conventions found [here](https://github.com/golang/go/wiki/CodeReviewComments).
* If in doubt, try to match your code to the current codebase.
* Create a pull request with a description of your changes.

I'll review pull requests as they come in and merge them if everything checks out.

Any and all contributions are greatly appreciated. Thank you!