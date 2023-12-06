package main

import (
	"os"
	"runtime/pprof"

	"github.com/tlj/aoc2023/aoc"
	"github.com/tlj/aoc2023/day06"
)

func main() {
	day := "day06"

	f, _ := os.Create(day + ".prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	d := aoc.NewDay("data/" + day + ".txt")
	d.Run(day06.Part1)
	d.Run(day06.Part2)
}
