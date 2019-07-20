package main

import "fmt"

func stringsJoin(sep string, strings ...string) string {
	joinedString := strings[0]

	for _, element := range strings[1:] {
		joinedString += sep
		joinedString += element
	}
	return joinedString
}

func main() {
	fmt.Printf("\"%s\"\n", stringsJoin(" ", "dog", "foobar", "fish"))
}
