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
	expected := []int64{ 467, 35, 633, 617, 592, 755, 664, 598 }

	actual := parts(example)

	assert.Equal(t, expected, actual)
}
