package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var example = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestDayThree(t *testing.T) {
	expectedNumbers := []int64{ 467, 35, 633, 617, 592, 755, 664, 598 }
	expectedRatios := []int64{ 16345, 451490 }

	actualNumbers, actualRatios := resolveSchematic(example)

	assert.Equal(t, expectedNumbers, actualNumbers)
	assert.Equal(t, expectedRatios, actualRatios)
}
