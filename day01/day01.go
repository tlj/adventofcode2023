package day01

import (
	"sort"

	"github.com/tlj/aoc2023/aoc"
)

func Part1(d aoc.Day) int {
	vals := d.LinesAsInts()
	var c int
	var m int
	for _, v := range vals {
		if v == nil {
			if c > m {
				m = c
			}
			c = 0
			continue
		}
		c += *v
	}
	if c > m {
		m = c
	}
	return m
}

func Part2(d aoc.Day) int {
	vals := d.LinesAsInts()
	var c int
	var m []int
	for _, v := range vals {
		if v == nil {
			if c > 0 {
				m = append(m, c)
			}
			c = 0
			continue
		}
		c += *v
	}
	if c > 0 {
		m = append(m, c)
	}
	var res int
	sort.Ints(m)
	largest := m[len(m)-3:]
	for _, l := range largest {
		res += l
	}
	return res
}
