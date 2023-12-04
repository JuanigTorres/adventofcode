package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/juanigtorres/adventofcode/slices"
)

//go:embed input.txt
var input string

var numbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func sorted(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return numbers[keys[i]] < numbers[keys[j]] 
	})

	return keys
}

func parse(s string) string {
	var first string
	min := len(s)

	for key := range numbers {
		pos := strings.Index(s, key)
		if pos != -1 && pos < min {
			min = pos
			first = key
		}
	}

	result := strings.ReplaceAll(s, first, numbers[first])
	for _, original := range sorted(numbers) {
		replacement := numbers[original]
		result = strings.ReplaceAll(result, original, replacement)
	}

	return result
}

func digits(s string) []int64 {
	result := make([]int64, 0, len(s))
	for _, value := range s {
		r, err := strconv.ParseInt(string(value), 10, 64)
		if err != nil {
			continue
		}
		result = append(result, r)
	}
	return result
}

func extremes(digits []int64) int64 {
	lenght := len(digits)
	if lenght == 0 {
		return 0
	}
	s := fmt.Sprintf("%d%d", digits[0], digits[lenght-1])
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

func total(lines []string) int64 {
	return slices.Reduce(lines, func(total int64, line string) int64 {
		p := parse(line)
		d := digits(p)
		v := extremes(d)
		return total + v
	}, 0)
}

func main() {
	lines := strings.Split(input, "\n")
	result := total(lines)
	fmt.Printf("result: %d\n", result)
}
