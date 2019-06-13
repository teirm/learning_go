// Fetch prints the content found at each specified URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if strings.HasPrefix(url, "http://") == false {
			url = "http://" + url
		}

		fmt.Printf("Getting url %s\n", url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copy %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
		fmt.Printf("HTTP Status: %s\n", resp.Status)
	}
}
