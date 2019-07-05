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

func main() {
   
    var weekOld []Issue
    var monthOld []Issue
    var yearOld  []Issue
    var ancient  []Issue

    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }

    for _, item := range result.Items {
        daysOld := daysAgo(item.CreatedAt)
        if daysOld <= 7 {
            append(weekOld, item)
        } else if daysOld > 7 && daysOld <= 30 {
            append(monthOld, item)
        } else if daysOld > 30 && daysOld <= 365 {
            append(yearOld, item)
        } else {
            append(ancient, item)
        }
    }

    fmt.Printf("%d issues:\n", result.TotalCount)
    for _, item := range result.Items {
        fmt.Printf("#%-5d %9.9s %.55s %d\n",
                item.Number, 
                item.User.Login, 
                item.Title, 
                daysAgo(item.CreatedAt))
    }
}
