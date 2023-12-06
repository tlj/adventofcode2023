package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"strings"
	"text/template"
)

type data struct {
	Day     string
	Profile bool
}

var targets = map[string]struct {
	target         string
	allowOverwrite bool
}{
	"templates/main.go.tmpl":     {target: "main.go", allowOverwrite: true},
	"templates/day.go.tmpl":      {target: "day{day}/day{day}.go"},
	"templates/day_test.go.tmpl": {target: "day{day}/day{day}_test.go"},
}

func writeFile(filename string, content []byte, overwrite bool) error {
	fmt.Printf("Writing %s... ", filename)
	_, err := os.Stat(filename)
	if err == nil && !overwrite {
		fmt.Println("Skipped.")
		return nil
	}

	err = os.WriteFile(filename, content, 0644)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("")

	return err
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: generate [daynum]")
		return
	}

	day := os.Args[1]

	// write empty data files if they don't already exist
	writeFile(fmt.Sprintf("data/day%s.txt", day), []byte{}, false)
	writeFile(fmt.Sprintf("data/day%sex.txt", day), []byte{}, false)

	// make directory
	os.Mkdir("day"+day, 0755)

	for source, target := range targets {
		filename := strings.ReplaceAll(target.target, "{day}", day)
		tmpl, err := os.ReadFile(source)
		if err != nil {
			panic(err)
		}

		t, err := template.New("day").Parse(string(tmpl))
		if err != nil {
			panic(err)
		}
		var content bytes.Buffer
		err = t.Execute(&content, data{Day: day, Profile: false})
		if err != nil {
			panic(err)
		}
		formatted, err := format.Source(content.Bytes())
		if err != nil {
			panic(err)
		}

		writeFile(filename, formatted, target.allowOverwrite)
	}
}
