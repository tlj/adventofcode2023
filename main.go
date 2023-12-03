package main

import (
	"github.com/tlj/aoc2023/aoc"
	"github.com/tlj/aoc2023/day03"
)

func main() {
	d := aoc.NewDay("data/day03.txt")
	d.Run(day03.Part1)
	d.Run(day03.Part2)
}
