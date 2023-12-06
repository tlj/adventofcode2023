package main

import (
	"github.com/tlj/aoc2023/aoc"
	"github.com/tlj/aoc2023/day06"
)

func main() {

	d := aoc.NewDay("data/day06.txt")
	d.Run(day06.Part1)
	d.Run(day06.Part2)
}
