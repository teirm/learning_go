// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		counts["stdin"] = make(map[string]int)
		countLines(os.Stdin, counts["stdin"])
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			counts[arg] = make(map[string]int)
			countLines(f, counts[arg])
			f.Close()
		}
	}
	for file_name, word_map := range counts {
		for line, n := range word_map {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", file_name, n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
