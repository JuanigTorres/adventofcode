package main

import (
	"math"
	"strings"

	"github.com/juanigtorres/adventofcode/slices"
)

type (
	Scratchcards []Scratchcard
	Scratchcard  struct {
		winners, numbers []int
	}
)

func NewScratchcardsFromString(s string) Scratchcards {
	rows := strings.Split(s, "\n")

	r := make(Scratchcards, 0, len(rows))
	for _, row := range rows {
		parts := strings.Split(strings.Split(row, ":")[1], "|")
		w := slices.Filter(strings.Split(parts[0], " "), isEmpty)
		n := slices.Filter(strings.Split(parts[1], " "), isEmpty)
		scratchcard := Scratchcard{
			winners: numbers(w),
			numbers: numbers(n),
		}
		r = append(r, scratchcard)
	}
	return r
}

func (s Scratchcard) Wins() int {
	wins := 0
	for _, winner := range s.winners {
		for _, number := range s.numbers {
			if number == winner {
				wins++
			}
		}
	}
	return wins
}

func (s Scratchcard) Points() int {
	count := s.Wins()
	if count > 0 {
		return int(math.Pow(float64(2), float64(count-1)))
	}
	return count
}

func (ss Scratchcards) TotalPoints() int {
	return slices.Reduce(ss, func(prev int, scratchcard Scratchcard) int {
		return prev + scratchcard.Points()
	}, 0)
}

func (ss Scratchcards) TotalCards() int {
	return totalCards(ss, 0)
}

func totalCards(ss Scratchcards, total int) int {
	for i, scratchcard := range ss {
		next := i + 1
		offset := scratchcard.Wins()
		copies := ss[next:next+offset]
		total = totalCards(copies, total + 1)
	}

	return total
}
