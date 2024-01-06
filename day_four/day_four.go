package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string


func main() {
	fmt.Println("Part 1", NewScratchcardsFromString(input).TotalPoints())
	fmt.Println("Part 2", NewScratchcardsFromString(input).TotalCards())
}
