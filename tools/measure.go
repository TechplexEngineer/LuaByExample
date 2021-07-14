package tools

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

var commentPat = regexp.MustCompile(`\s*(//|#|--)`)

const maxLineLength = 65

// Count the number of characters in each code line for
// all enabled examples. Print a message for each file that has lines over maxLineLength
// set the exit code to 1.
func Measure() {
	exampleNames := GetListOfExamples("examples.txt")

	foundLongLine := false

	for _, exampleName := range exampleNames {
		files := mustGlob("examples/" + BuildExampleId(exampleName) + "/*")
		for _, file := range files {
			lines := readLines(file)
			if lineNumber := CheckLongLines(lines); lineNumber != -1 {
				fmt.Printf("Line too long: %s:%d\n", file, lineNumber) //nolint:forbidigo

				foundLongLine = true
			}
		}
	}

	if foundLongLine {
		os.Exit(1)
	}
}

// Find the first line longer than maxLineLength and return the line number
// or -1 if no lines are too long.
func CheckLongLines(lines []string) int {
	for lineIndex, line := range lines {
		// Convert tabs to spaces before measuring, so we get an accurate measure
		// of how long the output will end up being.
		line := strings.ReplaceAll(line, "\t", "    ")
		if !commentPat.MatchString(line) && (utf8.RuneCountInString(line) > maxLineLength) {
			lineNumber := lineIndex + 1

			return lineNumber
		}
	}

	return -1
}
