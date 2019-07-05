package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/teirm/learning_go/ch_4/github"
)

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func printItem(item *github.Issue) {
	fmt.Printf("#%-5d %9.9s %.55s %d\n",
		item.Number,
		item.User.Login,
		item.Title,
		daysAgo(item.CreatedAt))
}

func main() {

	var weekOld []*github.Issue
	var monthOld []*github.Issue
	var yearOld []*github.Issue
	var ancient []*github.Issue

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Items {
		daysOld := daysAgo(item.CreatedAt)
		if daysOld <= 7 {
			weekOld = append(weekOld, item)
		} else if daysOld > 7 && daysOld <= 30 {
			monthOld = append(monthOld, item)
		} else if daysOld > 30 && daysOld <= 365 {
			yearOld = append(yearOld, item)
		} else {
			ancient = append(ancient, item)
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("--------------------------------")

	fmt.Println("Week-Old")
	for _, item := range weekOld {
		printItem(item)
	}

	fmt.Println("Month-Old")
	for _, item := range monthOld {
		printItem(item)
	}

	fmt.Println("Year-Old")
	for _, item := range yearOld {
		printItem(item)
	}

	fmt.Println("Ancient")
	for _, item := range ancient {
		printItem(item)
	}
}
