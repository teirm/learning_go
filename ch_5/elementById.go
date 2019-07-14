// ElementById calls the function pre(x) and post(x) for
// each node x in the tree rooted at n searching for the
// node matching the given ID.  Both are optional.

package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

func searchEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {

	var stopTraversal bool

	if pre != nil {
		stopTraversal = pre(n, id)
	}

	if stopTraversal {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := searchEachNode(c, id, pre, post)
		if result != nil {
			return result
		}
	}

	if post != nil {
		stopTraversal = post(n, id)
	}

	if stopTraversal {
		return n
	}

	return nil
}

func containsId(n *html.Node, id string) bool {

	for _, attr := range n.Attr {
		if attr.Key == "id" && attr.Val == id {
			return true
		}
	}
	return false
}

func ElementById(doc *html.Node, id string) *html.Node {

	var n = searchEachNode(doc, id, containsId, containsId)

	if n != nil {
		fmt.Printf("Found match: <%s>\n", n.Data)
	} else {
		fmt.Printf("No element matching id %s\n", id)
	}

	return n
}

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "elementById: %v\n", err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elementById: %v\n", err)
		os.Exit(1)
	}
	resp.Body.Close()
	ElementById(doc, "lowframe")
}
