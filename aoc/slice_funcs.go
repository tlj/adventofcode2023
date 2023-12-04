package aoc

import (
	"sort"
	"strconv"
	"strings"
)

func CreateList(a string) (c []string) {
	return strings.Split(a, " ")
}

func CreateListSorted(a string) (c []string) {
	c = CreateList(a)
	sort.Strings(c)
	return c
}

func CreateIntList(a string) (b []int) {
	spl := CreateList(a)
	for _, n := range spl {
		if n == "" {
			continue
		}
		c, err := strconv.Atoi(n)
		if err == nil {
			b = append(b, c)
		}
	}
	return b
}

func CreateIntListSorted(a string) (b []int) {
	list := CreateIntList(a)
	sort.Ints(list)
	return list
}

func IntersectSortedIntList(a, b []int) (c []int) {
	for _, v := range a {
		idx := sort.Search(len(b), func(i int) bool {
			return b[i] >= v
		})
		if idx < len(b) && b[idx] == v {
			c = append(c, v)
		}
	}
	return c
}
