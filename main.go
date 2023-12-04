package main

import (
	"github.com/tlj/aoc2023/aoc"
	"github.com/tlj/aoc2023/day04"
)

func main() {
	d := aoc.NewDay("data/day04.txt")
	d.Run(day04.Part1)
	d.Run(day04.Part2)
}
