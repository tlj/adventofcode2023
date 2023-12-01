package day01

import (
	"testing"

	"github.com/tlj/aoc2023/aoc"
)

func Test_Part1(t *testing.T) {
	d := aoc.NewDay("../data/day01ex.txt")
	res := Part1(d)
	expected := 142
	if res != expected {
		t.Errorf("Expected: %v. Got: %v.", expected, res)
	}
}

func Test_Part2(t *testing.T) {
	d := aoc.NewDay("../data/day01ex2.txt")
	res := Part2(d)
	expected := 281
	if res != expected {
		t.Errorf("Expected: %v. Got: %v.", expected, res)
	}
}
