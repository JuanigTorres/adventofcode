package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var example = []string{
	"two1nine",
	"eightwothree",
	"abcone2threexyz",
	"xtwone3four",
	"4nineeightseven2",
	"zoneight234",
	"7pqrstsixteen",
}

func TestDayOne(t *testing.T) {
	expected := int64(281)
	assert.Equal(t, expected, total(example))
}
