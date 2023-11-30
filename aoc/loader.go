package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Loader struct {
}

func LoadAsLines(input string) []string {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("\nLoaded %d lines.\n\n", len(lines))

	return lines
}

func ReadFileAsLines(inputFile string) []string {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	return LoadAsLines(string(content))
}
