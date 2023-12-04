package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/juanigtorres/adventofcode/slices"
)

//go:embed input.txt
var input string

func parse(game string) (string, map[string]int64) {
	result := map[string]int64{
		"red": 0,
		"green": 0,
		"blue": 0,
	}
	
	g := strings.Split(game, ": ")
	id, info := strings.Split(g[0], " ")[1], g[1]
	sets := strings.Split(info, "; ")
	for _, set := range sets {
		cubes := strings.Split(strings.TrimSpace(set), ",")
		for _, cube := range cubes {
			c := strings.Split(strings.TrimSpace(cube), " ")
			name := c[1]
			count, _ := strconv.ParseInt(c[0], 10, 64)
			if result[name] < count {
				result[name] = count
			}
		}
	}
	return id, result
}

func total(games []string, red, green, blue int64) (int64, int64) {
	gids := make([]string, 0, len(games))
	powers := make([]int64, 0, len(games))
	for _, game := range games {
		gid, r := parse(game)
		p := r["red"] * r["green"] * r["blue"]
		powers = append(powers, p)
		if r["red"] <= red && r["green"] <= green && r["blue"] <= blue {
			gids = append(gids, gid)
		}
	}
	
	t := slices.Reduce(gids, func(total int64, curr string) int64 {
		id, _ := strconv.ParseInt(curr, 10, 64)
		return total + id
	}, 0)

	p := slices.Reduce(powers, func(total int64, curr int64) int64 { 
		return total + curr
	}, 0)
	
	return t, p
}

func main() {
	games := strings.Split(input, "\n")
	part1, part2 := total(games, 12, 13, 14)
	fmt.Printf("results - Part 1: %d, Part 2: %d\n", part1, part2)
}