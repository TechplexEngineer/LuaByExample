package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"
)

var commentPat = regexp.MustCompile("\\s*\\/\\/")

const maxLineLength = 58

func Measure() {
	sourcePaths, err := filepath.Glob("./examples/*/*")
	check(err)
	foundLongFile := false
	for _, sourcePath := range sourcePaths {
		foundLongLine := false
		lines := readLines(sourcePath)
		for i, line := range lines {
			// Convert tabs to spaces before measuring, so we get an accurate measure
			// of how long the output will end up being.
			line := strings.Replace(line, "\t", "    ", -1)
			if !foundLongLine && !commentPat.MatchString(line) && (utf8.RuneCountInString(line) > maxLineLength) {
				fmt.Printf("measure: %s:%d\n", sourcePath, i+1)
				foundLongLine = true
				fmt.Printf("Found line that is %d chars, limit is %d, over by %d chars. Offending line in %s is:\n%s\n", utf8.RuneCountInString(line), maxLineLength, utf8.RuneCountInString(line)-maxLineLength, sourcePath, line)
				return
			}
		}
	}
	if foundLongFile {
		os.Exit(1)
	}
}
