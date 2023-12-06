package aoc

import (
	"bufio"
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

	return lines
}

func ReadFileAsLines(inputFile string) []string {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	return LoadAsLines(string(content))
}
