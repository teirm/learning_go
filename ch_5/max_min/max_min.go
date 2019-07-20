package main

import "fmt"

func max(vals ...int) int {

	if len(vals) == 0 {
		panic("No arguments")
	}

	m := vals[0]
	for _, val := range vals[1:] {
		if val > m {
			m = val
		}
	}
	return m
}

func min(vals ...int) int {
	if len(vals) == 0 {
		panic("No arguments")
	}

	m := vals[0]
	for _, val := range vals[1:] {
		if val < m {
			m = val
		}
	}
	return m
}

func main() {
	var max1 = max(1, 2, 3, 4, 5)
	fmt.Printf("%d\n", max1)

	var max2 = max()
	fmt.Printf("%d\n", max2)
}
