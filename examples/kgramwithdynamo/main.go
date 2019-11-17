package main

import (
	"fmt"

	"github.com/maisiesadler/search/examples/dynamoindex"
	"github.com/maisiesadler/search/index"
)

func main() {
	idx, _ := dynamoindex.CreateDynamoDictionary()
	kgramIdx := index.CreateKgramFromIndex(idx)
	tokens := make(chan string)
	go func() {
		defer close(tokens)
		tokens <- "weve"
		tokens <- "were"
		tokens <- "wont"
		tokens <- "while"
	}()
	kgramIdx.Add("doc-id", tokens)

	found, results := kgramIdx.Find("we*e")
	if found {
		fmt.Printf("Found %v result\n", len(results))
	} else {
		fmt.Printf("Did not find any results\n")
	}
	for _, result := range results {
		fmt.Printf("- '%v' = %v\n", result.Word, resultsToString(result.Matches))
	}
}

func resultsToString(m map[string]int) string {
	s := ""
	for k := range m {
		s += k + ","
	}

	return s
}
