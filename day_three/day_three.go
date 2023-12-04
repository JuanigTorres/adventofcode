package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/juanigtorres/adventofcode/slices"
)

//go:embed input.txt
var input string

type Number struct {
	value, start, end int
}

func matrix(s string) [][]rune {
	result := make([][]rune, 0)
	rows := strings.Split(s, "\n")

	for _, row := range rows {
		runes := bytes.Runes([]byte(row))
		result = append(result, runes)
	}

	return result
}

func allEmptySymbols(runes []rune) bool {
	hasSpecialChar := slices.Some(runes, func(r rune) bool {
		return r != '.'
	})
	return !hasSpecialChar
}

func isDigit(r rune) bool {
	digits := []rune{ '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	return slices.Some(digits, func(d rune) bool { return d == r })
}

func adyacents(row []rune) []Number {
	lenght := len(row)
	result := make([]Number, 0)

	for i := 0; i < lenght; i++ {
		if isDigit(row[i]) {
			start := i
			end := i
			for end < lenght - 1 && isDigit(row[end]) {
				end++
			}
			value, _ := strconv.Atoi(string(row[start:end]))
			n := Number{
				start: start,
				end:   end,
				value: value,
			}
			result = append(result, n)
			i = end
		}
	}

	return result
}

func neightbors(num Number, mtrx [][]rune, i int) []rune {
	result := make([]rune, 0)
	from, to := num.start, num.end

	if num.start > 0 {
		from--
		result = append(result, mtrx[i][from])
	}

	if num.end < len(mtrx[i]) {
		result = append(result, mtrx[i][to])
	}

	if i > 0 {
		result = append(result, mtrx[i-1][from:to+1]...)
	}

	if i < len(mtrx[i])-1 {
		result = append(result, mtrx[i+1][from:to+1]...)
	}

	return result
}

func parts(input string) []int64 {
	mtrx := matrix(input)
	result := make([]int64, 0)

	for i := 0; i < len(mtrx); i++ {
		nums := adyacents(mtrx[i])
		for _, num := range nums {
			if !allEmptySymbols(neightbors(num, mtrx, i)) {
				result = append(result, int64(num.value))
			}
		}
	}

	return result
}

func sum(numbers []int64) int64 {
	return slices.Reduce(numbers, func(prev, curr int64) int64 {
		return prev + curr
	}, 0)
}

func main() {
	total := sum(parts(input))
	fmt.Printf("total: %d\n", total)
}
