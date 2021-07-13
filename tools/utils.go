package tools

import (
	"io/ioutil"
	"strings"
)

// panic if err is not nil
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// read all lines from file at `path`
// returns slice of lines of file
// panics on error
func readLines(path string) []string {
	srcBytes, err := ioutil.ReadFile(path)
	check(err)
	return strings.Split(string(srcBytes), "\n")
}
