package main

import (
	"exercise/Chapter4/4_10/github"
	"fmt"
	"log"
	"os"
	"time"
)

type class string

const (
	mtoy class = "More than one year"
	ltoy class = "Less than one year"
	ltom class = "Less than one month"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	current := time.Now()
	classification := make(map[class]int)
	for _, issue := range result.Items {
		fmt.Println(issue.Title)
		if current.After(issue.CreatedAt.AddDate(1, 0, 0)) {
			classification[mtoy]++
		} else if current.Before(issue.CreatedAt.AddDate(0, 1, 0)) {
			classification[ltom]++
		} else {
			classification[ltoy]++
		}
	}
	fmt.Println("total results :", result.TotalCount)
	for class, cnt := range classification {
		fmt.Printf("%s : %d\n", class, cnt)
	}
}
