package worker

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Result struct {
	Line    string
	LineNum int
	Path    string
}

type Results struct {
	Inner []Result
}

func NewResult(line string, lineNum int, path string) Result {
	return Result{line, lineNum, path}
}

func FindInFile(path string, find string) *Results {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("unable to open the file: ", err)
	}
	results := Results{make([]Result, 0)}
	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, find) {
			r := NewResult(line, lineNum, path)
			results.Inner = append(results.Inner, r)
		}
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(results.Inner) == 0 {
		return nil
	} else {
		return &results
	}
}
