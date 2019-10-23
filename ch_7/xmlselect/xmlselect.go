package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement // stack of start elements

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				elements := elementsToStrings(stack)
				fmt.Printf("%s: %s\n", strings.Join(elements, " "), tok)
			}
		}
	}
}

// elementsToStrings converts an array of start elements to strings
// with attributes properly ordered
func elementsToStrings(x []xml.StartElement) []string {
	var elements []string

	// append token name
	for _, token := range x {
		elements = append(elements, token.Name.Local)
		// append all attrs right after token name
		for _, attr := range token.Attr {
			elements = append(elements, attr.Value)
		}
	}
	return elements
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []xml.StartElement, y []string) bool {

	elements := elementsToStrings(x)

	for len(y) <= len(elements) {
		if len(y) == 0 {
			return true
		}
		if elements[0] == y[0] {
			y = y[1:]
		}
		elements = elements[1:]
	}
	return false
}
