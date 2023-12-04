package day04

import (
	"fmt"
	"math"
	"strings"

	"github.com/tlj/aoc2023/aoc"
)

func Part1(d aoc.Day) int {
	sum := 0

	for _, l := range d.Lines {
		cs := strings.Split(l, ":")
		numss := strings.Split(string(cs[1]), "|")

		correct := aoc.CreateIntListSorted(numss[0])
		mine := aoc.CreateIntListSorted(numss[1])

		winning := aoc.IntersectSortedList(correct, mine)
		sum += int(math.Floor(math.Pow(2, float64(len(winning))) / 2))
	}

	return sum
}

func Part2(d aoc.Day) int {
	sum := len(d.Lines)

	for ln := 0; ln < len(d.Lines); ln++ {
		l := d.Lines[ln]

		cs := strings.Split(l, ":")
		numss := strings.Split(string(cs[1]), "|")

		correct := aoc.CreateIntListSorted(numss[0])
		mine := aoc.CreateIntListSorted(numss[1])

		var cardnum int
		fmt.Sscanf(l, "Card %d:", &cardnum)

		lsum := 0
		winning := aoc.IntersectSortedList(correct, mine)
		for _ = range winning {
			lsum++

			d.Lines = append(d.Lines, d.Lines[cardnum-1+lsum])
		}
		sum += lsum
	}

	return sum
}
