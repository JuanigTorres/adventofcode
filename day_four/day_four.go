package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string


func main() {
	total := NewScratchcardsFromString(input).TotalPoints()
	fmt.Println("Total", total)
}
