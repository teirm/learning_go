// forEachNode calls the function pre(x) and post(x) for each node
// x in the tree rooted at n.  Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder)

package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild != nil {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		} else {
			fmt.Printf("%*s<%s\\>\n", depth*2, "", n.Data)
		}
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func main() {

	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline2: %v\n", err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline2: %v\n", err)
		os.Exit(1)
	}
	resp.Body.Close()
	forEachNode(doc, startElement, endElement)
}
