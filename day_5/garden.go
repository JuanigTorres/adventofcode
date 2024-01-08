package main

import (
	"github.com/juanigtorres/adventofcode/functional"
	"github.com/juanigtorres/adventofcode/slices"
	"github.com/juanigtorres/adventofcode/strings"
)

type Seeds []string

func CreateSeedsromInput(input string) Seeds {
	GetFirstValue := functional.InverseCurry2(slices.Value[[]string, string])(0)
	DeleteSeedsSubString := functional.InverseCurry2(strings.DeleteSubstring)("seeds:")

	return functional.Pipe4(
		strings.Lines,
		GetFirstValue,
		DeleteSeedsSubString,
		strings.Words,
	)(input)
}

type Almanac struct {
	from, to           string
	start, end, lenght int
}

func CreateAlmanacFromInput(input string) Almanac {
	return Almanac{}
}
