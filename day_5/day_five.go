package main

import (
	"fmt"

	"github.com/juanigtorres/adventofcode/resources"
)

var input string = resources.InputStringByDay(5)

func main() {
	fmt.Println("Part 1", CreateSeedsromInput(input))
}
