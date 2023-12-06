package day06

import (
	"strconv"
	"strings"

	"github.com/tlj/aoc2023/aoc"
)

func Part1(d aoc.Day) int {
	times := d.IntsInLine(d.Lines[0])
	distances := d.IntsInLine(d.Lines[1])

	sum := 1
	for k, t := range times {
		wins := 0

		for i := 0; i < t; i++ {
			if (t-i)*i > distances[k] {
				wins++
			}
		}

		sum *= wins
	}

	return sum
}

func Part2(d aoc.Day) int {
	sum := 0

	timeSpl := strings.Split(d.Lines[0], ":")
	timeStr := strings.ReplaceAll(timeSpl[1], " ", "")
	t, _ := strconv.Atoi(timeStr)

	distSpl := strings.Split(d.Lines[1], ":")
	distStr := strings.ReplaceAll(distSpl[1], " ", "")
	dist, _ := strconv.Atoi(distStr)

	for i := 0; i < t; i++ {
		if (t-i)*i > dist {
			sum++
		}
	}

	return sum
}
