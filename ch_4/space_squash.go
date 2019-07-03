// condense any unicode spaces to a single ascii space in-place
package main

import (
	"fmt"
	"unicode"
)

func squashSpaces(bytes []byte) []byte {

	spaceIndex := 0
	spaceFlag := false

	for _, s := range bytes {
		fmt.Printf("space_flag %t\n", spaceFlag)
		if unicode.IsSpace(rune(s)) {
			if !spaceFlag {
				bytes[spaceIndex] = ' '
				spaceFlag = true
				spaceIndex++
			}
		} else {
			bytes[spaceIndex] = s
			spaceFlag = false
			spaceIndex++
		}
	}

	return bytes[:spaceIndex]
}

func main() {

	s := []byte("a   b    cd   e")

	s = squashSpaces(s)

	fmt.Println(string(s))
}
