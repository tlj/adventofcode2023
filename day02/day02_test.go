package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tlj/aoc2023/aoc"
)

func Test_Part1(t *testing.T) {
	d := aoc.NewDay("../data/day02ex.txt")

	games := parseInputToGames(d.Lines)
	if len(games) != 5 {
		t.Errorf("Expected: 5 games. Got: %v.", len(games))
	}

	assert.Equalf(t, game{
		number:        1,
		valid:         true,
		minNumOfCubes: map[string]int{"blue": 6, "red": 4, "green": 2},
		turns: []turn{
			{cubes: map[string]int{"blue": 3, "red": 4}},
			{cubes: map[string]int{"red": 1, "green": 2, "blue": 6}},
			{cubes: map[string]int{"green": 2}},
		},
	}, games[0], "First game is same.")
	assert.Equalf(t, 48, games[0].power(), "First game power is same.")

	res := Part1(d)
	expected := 8
	if res != expected {
		t.Errorf("Expected: %v. Got: %v.", expected, res)
	}
}

func Test_Part2(t *testing.T) {
	d := aoc.NewDay("../data/day02ex.txt")
	res := Part2(d)
	expected := 2286
	if res != expected {
		t.Errorf("Expected: %v. Got: %v.", expected, res)
	}
}
