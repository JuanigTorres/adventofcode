package main

import (
	"strconv"
	"strings"

	"github.com/juanigtorres/adventofcode/slices"
)

func numbers(s []string) []int {
	return slices.Map(s, func(n string) int {
		v, _ := strconv.ParseInt(strings.TrimSpace(n), 10, 64)
		return int(v)
	})
}

func isEmpty(s string) bool {
	return len(s) == 0
}
