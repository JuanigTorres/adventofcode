package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var example = []string{
	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}

func TestDayTwo(t *testing.T) {

	expectedPart1 := int64(8)
	expectedPart2 := int64(2286)

	actualPart1, actualPart2 := total(example, 12, 13, 14)

	assert.Equal(t, expectedPart1, actualPart1)
	assert.Equal(t, expectedPart2, actualPart2)
}
