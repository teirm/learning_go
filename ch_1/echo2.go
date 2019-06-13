package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Printf("Index: %d\tValue: %s\n", index, arg)
	}
}
