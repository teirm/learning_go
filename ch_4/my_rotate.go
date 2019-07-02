package main

import "fmt"

func rotateLeft(s []int, n int) []int {

	if n > len(s) {
		n = n % len(s)
	}

	a := s[:n]
	b := s[n:]

	return append(b, a...)
}

func rotateRight(s []int, n int) []int {
	if n > len(s) {
		n = n % len(s)
	}

	shift := len(s) - n

	a := s[:shift]
	b := s[shift:]

	return append(b, a...)
}

func main() {

	s := []int{0, 1, 2, 3, 4, 5}

	left := rotateLeft(s, 7)
	fmt.Println(left)

	right := rotateRight(s, 2)
	fmt.Println(right)
}
