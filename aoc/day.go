package aoc

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day struct {
	Filename string
	Data     []byte
	Lines    []string
}

func NewDay(name string) Day {
	d := Day{
		Filename: name,
	}

	var err error
	d.Data, err = os.ReadFile(d.Filename)
	if err != nil {
		panic(err)
	}

	d.Lines = ReadFileAsLines(d.Filename)
	if err != nil {
		panic(err)
	}

	return d
}

func (d Day) IntsInLine(l string) (ints []int) {
	spl := strings.Split(l, " ")

	for _, ts := range spl {
		t, err := strconv.Atoi(ts)
		if err == nil {
			ints = append(ints, t)
		}
	}

	return ints
}

func (d Day) LinesAsInts() []*int {
	var res []*int
	for _, l := range d.Lines {
		if strings.TrimSpace(l) == "" {
			res = append(res, nil)
			continue
		}
		v, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		res = append(res, &v)
	}

	return res
}

func (d Day) Run(fn func(d Day) int) {
	res := fn(d)
	fmt.Printf("Result: %d\n", res)
}
