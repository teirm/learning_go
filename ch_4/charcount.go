package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	typeHisto := make(map[string]int)
	var utflen [utf8.UTFMax + 1]int

	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsLetter(r) {
			typeHisto["letter"]++
		}
		if unicode.IsNumber(r) {
			typeHisto["number"]++
		}
		if unicode.IsSpace(r) {
			typeHisto["space"]++
		}
		if unicode.IsPunct(r) {
			typeHisto["punct"]++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Println("rune\tcounts\t")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	fmt.Print("\ntype\tcount\n")
	for t, n := range typeHisto {
		fmt.Printf("%s\t%d\n", t, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
