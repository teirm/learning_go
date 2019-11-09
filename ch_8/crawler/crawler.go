// crawler is a concurrent web crawler
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/teirm/learning_go/ch_5/links"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests

// TODO(teirm) notes on rate limiting -- can have each
// link be in a struct with depth -- don't add to
// worklist if depth is greater than desired depth OR
// don't spawn a goroutine if greater than desired depth

// need some function to convert []string into []CrawlerTask
// func createCrawlerTask(url string, depth int) CrawlerTask
// func createCrawlerTasks(url []string, depth int) []CrawlerTask

type CrawlerTask struct {
	url   string
	depth int
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist
	var depth = flag.Int("depth", 10, "maximum dept of URL crawl")

	flag.Parse()
	fmt.Println(depth)

	// start with the command-line arguments
	n++
	go func() { worklist <- os.Args[3:] }()

	// Crawl the web concurrently
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
