package day04

import (
	"testing"

	"github.com/tlj/aoc2023/aoc"
)

func Test_Part1(t *testing.T) {
	d := aoc.NewDay("../data/day04ex.txt")
	res := Part1(d)
	expected := 13
	if res != expected {
		t.Errorf("Expected: %v. Got: %v.", expected, res)
	}
}

func Test_Part2(t *testing.T) {
	d := aoc.NewDay("../data/day04ex.txt")
	res := Part2(d)
	expected := 30
	if res != expected {
		t.Errorf("Expected: %v. Got: %v.", expected, res)
	}
}
