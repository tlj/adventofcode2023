package main

import (
	"github.com/tlj/aoc2023/aoc"
	"github.com/tlj/aoc2023/day01"
)

func main() {
	d := aoc.NewDay("data/day01.txt")
	d.Run(day01.Part1)
	d.Run(day01.Part2)
}
