// Modify dup2 to print the names of all files in which each duplicated line occurs

package main

import (
	"bufio"
	"fmt"
	"os"
)

type holder struct {
	fileName []string
	count    int
}

func main() {
	counts := make(map[string]holder)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%s %d \t%s\n", n.fileName, n.count, line)
		}
	}
}

func countLines(f *os.File, counts map[string]holder) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		currentCount := counts[input.Text()]
		newCount := currentCount.count + 1
		newCurrentFileName := currentCount.fileName
		if !contains(f.Name(), newCurrentFileName) {
			newCurrentFileName = append(newCurrentFileName, f.Name())
		}
		counts[input.Text()] = holder{fileName: newCurrentFileName, count: newCount}
	}
	// NOTE: ignoring potential errors from input.Err()
}

func contains(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}
