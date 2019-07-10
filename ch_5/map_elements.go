package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {

	elementMap := make(map[string]int)

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for element, count := range histogram(elementMap, doc) {
		fmt.Printf("%s\t%d\n", element, count)
	}
}

func histogram(elementMap map[string]int, n *html.Node) map[string]int {

	if n == nil {
		return elementMap
	}

	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			elementMap["a"]++
		case "p":
			elementMap["p"]++
		case "meta":
			elementMap["meta"]++
		case "link":
			elementMap["link"]++
		case "div":
			elementMap["div"]++
		case "li":
			elementMap["li"]++
		case "form":
			elementMap["form"]++
		case "input":
			elementMap["input"]++
		case "button":
			elementMap["button"]++
		case "body":
			elementMap["body"]++
		case "ul":
			elementMap["ul"]++
		case "script":
			elementMap["script"]++
		case "section":
			elementMap["section"]++
		case "h1":
			elementMap["h1"]++
		default:
			elementMap["unknown"]++
		}

	}

	c := n.FirstChild
	elementMap = histogram(elementMap, c)
	elementMap = histogram(elementMap, n.NextSibling)

	return elementMap
}
