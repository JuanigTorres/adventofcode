package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/juanigtorres/adventofcode/resources"
	"github.com/juanigtorres/adventofcode/slices"
)

var input string = resources.InputStringByDay(3)

type Number struct {
	value, start, end int
	neightbors        []Symbol
}

func numbers(mtrx [][]rune, pos int) []Number {
	row := mtrx[pos]
	lenght := len(row)
	result := make([]Number, 0)

	for i := 0; i < lenght; i++ {
		if isDigit(row[i]) {
			start := i
			end := i + 1
			for end < lenght && isDigit(row[end]) {
				end++
			}
			value, _ := strconv.Atoi(string(row[start:end]))
			n := Number{
				start:      start,
				end:        end,
				value:      value,
				neightbors: symbols(value, start, end, mtrx, pos),
			}
			result = append(result, n)
			i = end
		}
	}

	return result
}

type (
	Point struct {
		x, y int
	}

	Symbol struct {
		value string
		point Point
	}
)

func symbols(num, start, end int, mtrx [][]rune, i int) []Symbol {
	result := make([]Symbol, 0)
	from, to := start, end

	if from > 0 {
		from--
		s := Symbol{
			value: string(mtrx[i][from]),
			point: Point{x: i, y: from},
		}
		result = append(result, s)
	}

	if to < len(mtrx[i]) {
		s := Symbol{
			value: string(mtrx[i][to]),
			point: Point{x: i, y: to},
		}
		result = append(result, s)
	}

	if i > 0 {
		x := i - 1
		below := make([]rune, 0)

		if to == len(mtrx[i]) {
			below = append(below, mtrx[x][from:to]...)
		}

		if to < len(mtrx[i]) {
			below = append(below, mtrx[x][from:to+1]...)
		}

		for idx, v := range below {
			y := idx + from
			s := Symbol{
				value: string(v),
				point: Point{x: x, y: y},
			}
			result = append(result, s)
		}
	}

	if i < len(mtrx[i])-1 {
		x := i + 1
		above := make([]rune, 0)

		if to == len(mtrx[i]) {
			above = append(above, mtrx[x][from:to]...)
		}

		if to < len(mtrx[i]) {
			above = append(above, mtrx[x][from:to+1]...)
		}

		for idx, v := range above {
			y := idx + from
			s := Symbol{
				value: string(v),
				point: Point{x: x, y: y},
			}
			result = append(result, s)
		}
	}

	return result
}

func gearRatios(numbers []Number) []int64 {
	ratios := make([]int64, 0)
	gears := make(map[string][]Number)

	for _, n := range numbers {
		for _, symbol := range n.neightbors {
			if symbol.value == "*" {
				k := fmt.Sprintf("(%d,%d)", symbol.point.x, symbol.point.y)
				gears[k] = append(gears[k], n)
			}
		}
	}

	for _, nums := range gears {
		if len(nums) == 2 {
			ratio := nums[0].value * nums[1].value
			ratios = append(ratios, int64(ratio))
		}
	}

	return ratios
}

func parseSchematic(s string) [][]rune {
	result := make([][]rune, 0)
	rows := strings.Split(s, "\n")

	for _, row := range rows {
		runes := bytes.Runes([]byte(row))
		result = append(result, runes)
	}

	return result
}

func parts(nums []Number) []int64 {
	result := make([]int64, 0)
	for _, num := range nums {
		if ok := slices.Some(num.neightbors, func(s Symbol) bool { return s.value != "." }); ok {
			result = append(result, int64(num.value))
		}
	}

	return result
}

func isDigit(r rune) bool {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	return slices.Some(digits, func(d rune) bool { return d == r })
}

func resolveSchematic(input string) ([]int64, []int64) {
	mtrx := parseSchematic(input)
	nums := make([]Number, 0)

	for i := 0; i < len(mtrx); i++ {
		nums = append(nums, numbers(mtrx, i)...)
	}

	return parts(nums), gearRatios(nums)
}

func sum(numbers []int64) int64 {
	return slices.Reduce(numbers, func(prev, curr int64) int64 {
		return prev + curr
	}, 0)
}

func main() {
	numbersParts, gearRatios := resolveSchematic(input)

	fmt.Printf("Part 1: %d\n", sum(numbersParts))
	fmt.Printf("Part 2: %d\n", sum(gearRatios))
}
