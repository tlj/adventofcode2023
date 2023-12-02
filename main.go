package main

import (
	"github.com/tlj/aoc2023/aoc"
	"github.com/tlj/aoc2023/day02"
)

func main() {
	d := aoc.NewDay("data/day02.txt")
	d.Run(day02.Part1)
	d.Run(day02.Part2)
}
