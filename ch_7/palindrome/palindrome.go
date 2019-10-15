package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {

	var middle int

	if s.Len()%2 == 0 {
		middle = s.Len() / 2
	} else {
		middle = (s.Len() / 2) + 1
	}

	for i, j := 0, s.Len()-1; i < middle; i, j = i+1, j-1 {
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
	}
	return true
}

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {

	a := []int{1, 2, 3, 2, 1}
	b := []int{1, 2, 3, 3, 2, 1}
	c := []int{1, 2, 3, 4, 5, 6}
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	cases := [][]int{a, b, c, d}

	var res bool
	for _, array := range cases {
		res = IsPalindrome(IntSlice(array))
		fmt.Printf("Result: %t\n", res)
	}

}
