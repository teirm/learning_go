package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column width and print table
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("3m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type customSort struct {
	t           []*Track
	less        func(x, y *Track, comps []func(x, y *Track) bool) bool
	comparators []func(x, y *Track) bool // slice of comparison functions
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j], x.comparators) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

// append a comparison function to the slice of comparators
func (x *customSort) AddComparator(comp func(x, y *Track) bool) {
	x.comparators = append(x.comparators, comp)
}

// clear the slice of comparators
func (x *customSort) ClearComparators() {
	x.comparators = make([]func(x, y *Track) bool, 0)
}

func main() {
	fmt.Println("Sort by Artist")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("Custom Sort")

	// create a custom sort
	var c = customSort{tracks, func(x, y *Track, comparators []func(x, y *Track) bool) bool {
		for _, comp := range comparators {
			result := comp(x, y)
			if result {
				break
			}
		}
		return false
	}, make([]func(x, y *Track) bool, 0)}

	c.AddComparator(func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		return false
	})

	c.AddComparator(func(x, y *Track) bool {
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		return false
	})

	c.AddComparator(func(x, y *Track) bool {
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	})

	sort.Sort(c)
	printTracks(tracks)
}
