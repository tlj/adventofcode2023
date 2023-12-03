package day03

import (
	"math"
	"regexp"
	"strconv"

	"github.com/tlj/aoc2023/aoc"
)

var (
	numEx = regexp.MustCompile(`[0-9]+`)
	astEx = regexp.MustCompile(`\*`)
)

func numpsInLines(lines []string) (ret [][][]int) {
	for _, l := range lines {
		ret = append(ret, numEx.FindAllStringIndex(l, -1))
	}
	return ret
}

func Part1(d aoc.Day) int {
	sum := 0
	for ln, numps := range numpsInLines(d.Lines) {
		l := d.Lines[ln]
		for _, ps := range numps {
			numstr := l[ps[0]:ps[1]]
			num, _ := strconv.Atoi(numstr)

			pl := ""
			var bab []int
			if ln > 0 {
				bab = append(bab, ln-1)
			}
			if ln < len(d.Lines)-1 {
				bab = append(bab, ln+1)
			}

			if ps[0] > 1 {
				pl += string(l[ps[0]-1])
			}
			if ps[1] < len(l)-1 {
				pl += string(l[ps[1]])
			}

			for _, b := range bab {
				pl += string(d.Lines[b])[int(math.Max(float64(ps[0]-1), 0)):int(math.Min(float64(ps[1]+1), float64(len(d.Lines[b]))))]
			}

			for j := 0; j < len(pl); j++ {
				if pl[j] != 46 {
					sum += num
					break
				}
			}
		}
	}
	return sum
}

func adjacentToNumps(lines []string, ln, x, x2 int, numps [][][]int) []int {
	var res []int
	for nln, nps := range numps {
		if math.Abs(float64(ln-nln)) <= 1 {
			for _, np := range nps {
				if x >= np[0]-1 && x2 <= np[1]+1 {
					numstr := lines[nln][np[0]:np[1]]
					num, _ := strconv.Atoi(numstr)
					res = append(res, num)
				}
			}
		}
	}
	return res
}

func Part2(d aoc.Day) int {
	sum := 0
	numps := numpsInLines(d.Lines)
	for ln, l := range d.Lines {
		asts := astEx.FindAllStringIndex(l, -1)
		for _, ast := range asts {
			adj := adjacentToNumps(d.Lines, ln, ast[0], ast[1], numps)
			if len(adj) == 2 {
				sum += adj[0] * adj[1]
			}
		}
	}
	return sum
}
