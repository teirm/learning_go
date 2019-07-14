package main

import (
	"fmt"
	"os"
	"strings"
)

func expand(s string, sub string, f func(string) string) string {

	replacement := f(sub)

	return strings.ReplaceAll(s, sub, replacement)

}

func addDog(s string) string {
	return s + "dog"
}

func main() {
	initial := os.Args[1]
	subString := os.Args[2]

	fmt.Printf("initial: %s substring:%s\n", initial, subString)

	final := expand(initial, subString, addDog)

	fmt.Printf("final: %s\n", final)
}
