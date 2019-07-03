package main

import "fmt"

func removeAdjacents(s []string) []string {

	endIndex := 1
	currentAdj := s[0]

	for i := 1; i < len(s); i++ {
		if currentAdj != s[i] {
			s[endIndex] = s[i]
			currentAdj = s[i]
			endIndex++
		}
	}

	return s[:endIndex]
}

func main() {

	s := []string{"a", "a", "a", "b", "b", "a", "c", "a", "d", "a"}

	s = removeAdjacents(s)

	fmt.Println(s)
}
