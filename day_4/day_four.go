package main

import (
	"fmt"

	"github.com/juanigtorres/adventofcode/resources"
)

var input string = resources.InputStringByDay(4)

func main() {
	fmt.Println("Part 1", NewScratchcardsFromString(input).TotalPoints())
	fmt.Println("Part 2", NewScratchcardsFromString(input).TotalCards())
}
