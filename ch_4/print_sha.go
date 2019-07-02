package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {

	shaType := flag.Int("sha", 256, "sha size")

	flag.Parse()

	fmt.Printf("Sha type is %d\n", *shaType)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		switch *shaType {
		case 256:
			fmt.Printf("%x", sha256.Sum256([]byte(input.Text())))
		case 384:
			fmt.Printf("%x", sha512.Sum384([]byte(input.Text())))
		case 512:
			fmt.Printf("%x", sha512.Sum512([]byte(input.Text())))
		default:
			fmt.Fprintf(os.Stderr, "Invalid sha type: %d\n", *shaType)
			os.Exit(1)
		}
	}
}
