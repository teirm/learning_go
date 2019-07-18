package main

import (
	"fmt"
	"sort"
)

// prereqs map computer science course to their
// prerequisites
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

func topoSortMap(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(item string)

	visitAll = func(item string) {
		if !seen[item] {
			order = append(order, item)
		}
		for _, entry := range m[item] {
			if !seen[entry] {
				seen[entry] = true
				visitAll(entry)
				order = append(order, entry)
			}
		}
	}

	for key := range m {
		visitAll(key)
	}
	return order
}

func main() {
	fmt.Printf("Regular TopoSort\n")
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d\t%s\n", i+1, course)
	}

	fmt.Printf("Map based TopoSort\n")
	for i, course := range topoSortMap(prereqs) {
		fmt.Printf("%d\t%s\n", i+1, course)
	}
}
