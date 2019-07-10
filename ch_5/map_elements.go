package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {

	element_map := make(map[string]int)

	doc, err := html.Parse(os.stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for element, count := range histogram(element_map, doc) {
		fmt.Printf("%s\t%d\n", element, count)
	}
}
