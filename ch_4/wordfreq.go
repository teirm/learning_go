// count word frequencies in stdin
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	histogram := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		histogram[word]++
	}

	fmt.Println("word\t\t\tcount\n")
	for w, n := range histogram {
		fmt.Printf("%s\t\t\t%d\n", w, n)
	}
}
