package main

import (
	"os"
	"runtime/pprof"

	"github.com/tlj/aoc2023/aoc"
	"github.com/tlj/aoc2023/day05"
)

func main() {
	day := "day05"

	f, _ := os.Create(day + ".prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	d := aoc.NewDay("data/" + day + ".txt")
	d.Run(day05.Part1)
	d.Run(day05.Part2)
}
