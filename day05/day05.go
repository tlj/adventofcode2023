package day05

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/tlj/aoc2023/aoc"
)

func ParseInput(d aoc.Day) (seeds []int, maps [][][3]int) {
	for i := 0; i < len(d.Lines); i++ {
		l := d.Lines[i]

		if l == "" {
			continue
		}

		if strings.HasPrefix(l, "seeds: ") {
			l = strings.TrimPrefix(l, "seeds: ")
			seedStrs := strings.Split(l, " ")
			for _, ss := range seedStrs {
				si, _ := strconv.Atoi(ss)
				seeds = append(seeds, si)
			}
			continue
		}

		if strings.Contains(l, "map:") {
			i++
			var mapValues [][3]int
			for i < len(d.Lines) {
				l = d.Lines[i]
				if l == "" {
					break
				}
				i++
				var mapValue [3]int
				fmt.Sscanf(l, "%d %d %d", &mapValue[0], &mapValue[1], &mapValue[2])
				mapValues = append(mapValues, mapValue)
			}
			maps = append(maps, mapValues)
		}
	}

	return seeds, maps
}

func findLocationForSeed(seed int, maps [][][3]int) int {
	source := seed
	target := source

	for _, m := range maps {
		source = target
		for _, v := range m {
			if source >= v[1] && source < v[1]+v[2] {
				target = v[0] + (source - v[1])
				break
			}
		}
	}

	return target
}

func Part1(d aoc.Day) int {
	seeds, maps := ParseInput(d)

	lowest := math.MaxInt

	for _, s := range seeds {
		target := findLocationForSeed(s, maps)

		if target < lowest {
			lowest = target
		}
	}

	return lowest
}

func Part2(d aoc.Day) int {
	seeds, maps := ParseInput(d)

	lowest := math.MaxInt

	for sn := 0; sn < len(seeds); sn++ {
		for sj := seeds[sn]; sj < seeds[sn]+seeds[sn+1]; sj++ {
			target := findLocationForSeed(sj, maps)

			if target < lowest {
				lowest = target
			}
		}
		sn++
	}

	return lowest
}
