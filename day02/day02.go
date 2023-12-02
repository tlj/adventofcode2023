package day02

import (
	"fmt"
	"strings"

	"github.com/tlj/aoc2023/aoc"
)

type game struct {
	number        int
	valid         bool
	minNumOfCubes map[string]int
	turns         []turn
}

func (g *game) power() int {
	sum := 1
	for _, v := range g.minNumOfCubes {
		sum *= v
	}
	return sum
}

type turn struct {
	cubes map[string]int
}

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func (t turn) validate() bool {
	for k, v := range t.cubes {
		if v > maxCubes[k] {
			return false
		}
	}
	return true
}

func parseInputToGames(in []string) []game {
	var games []game
	for _, l := range in {
		gameRestSplit := strings.Split(l, ":")

		g := game{valid: true, minNumOfCubes: make(map[string]int)}
		fmt.Sscanf(gameRestSplit[0], "Game %d", &g.number)

		turnSplit := strings.Split(gameRestSplit[1], ";")
		for _, ts := range turnSplit {
			t := turn{cubes: make(map[string]int)}

			cubeSplit := strings.Split(ts, ",")
			for _, cs := range cubeSplit {
				var color string
				var throws int

				fmt.Sscanf(cs, "%d %s", &throws, &color)
				t.cubes[color] = throws

				if g.minNumOfCubes[color] < throws {
					g.minNumOfCubes[color] = throws
				}
			}
			g.turns = append(g.turns, t)
		}

		games = append(games, g)
	}

	return games
}

func Part1(d aoc.Day) int {
	games := parseInputToGames(d.Lines)
	invalidSum := 0
	for _, g := range games {
		for _, t := range g.turns {
			if !t.validate() {
				g.valid = false
				break
			}
		}
		if g.valid {
			invalidSum += g.number
		}
	}
	return invalidSum
}

func Part2(d aoc.Day) int {
	games := parseInputToGames(d.Lines)
	result := 0

	for _, g := range games {
		result += g.power()
	}

	return result
}
