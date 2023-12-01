package day01

import (
	"fmt"
	"strconv"

	"github.com/tlj/aoc2023/aoc"
)

func Part1(d aoc.Day) int {
	sum := 0
	for _, l := range d.Lines {
		var ints []int
		for i := 0; i < len(l); i++ {
			v, err := strconv.Atoi(string(l[i]))
			if err == nil {
				ints = append(ints, v)
			}
		}
		ss := fmt.Sprintf("%d%d", ints[0], ints[len(ints)-1])
		s, _ := strconv.Atoi(ss)
		sum += s
	}
	return sum
}

func Part2(d aoc.Day) int {
	nums := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	sum := 0
	for _, l := range d.Lines {
		var ints []int
		for i := 0; i < len(l); i++ {
			cur := string(l[i])
			for k, n := range nums {
				if i+len(k) > len(l) {
					continue
				}
				if l[i:i+len(k)] == k {
					cur = n
					break
				}
			}

			v, err := strconv.Atoi(cur)
			if err == nil {
				ints = append(ints, v)
			}
		}
		ss := fmt.Sprintf("%d%d", ints[0], ints[len(ints)-1])
		s, _ := strconv.Atoi(ss)
		sum += s
	}
	return sum
}
