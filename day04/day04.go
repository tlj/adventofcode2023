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

	parsedLines := make(map[int][]int, len(d.Lines))
	var cards []int

	for ln := 0; ln < len(d.Lines); ln++ {
		l := d.Lines[ln]

		var cardnum int
		fmt.Sscanf(l, "Card %d:", &cardnum)

		cards = append(cards, cardnum)

		cs := strings.Split(l, ":")
		numss := strings.Split(string(cs[1]), "|")

		correct := aoc.CreateIntListSorted(numss[0])
		mine := aoc.CreateIntListSorted(numss[1])
		winning := aoc.IntersectSortedList(correct, mine)

		parsedLines[cardnum] = winning
	}

	for i := 0; i < len(cards); i++ {
		ln := cards[i]
		lsum := 0

		for range parsedLines[ln] {
			lsum++

			cards = append(cards, ln+lsum)
		}

		sum += lsum
	}

	return sum
}
