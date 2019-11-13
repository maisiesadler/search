package main

import (
	"fmt"

	"github.com/maisiesadler/search/index"
	"github.com/maisiesadler/search/preprocess"
	"github.com/maisiesadler/search/tokenise"
)

func main() {
	doc1 := make(chan string)
	go func() {
		defer close(doc1)
		doc1 <- "hello this is some information in a document"
	}()

	tokens := tokenise.Tokenise(doc1)
	preprocessed := preprocess.Lowercase().Process(tokens)

	idx := index.Create()
	idx.Add("doc-1", preprocessed)

	found, results := idx.Find("information")
	fmt.Printf("Found %v", found)
	for _, result := range results {
		fmt.Println("Document - " + toString(result.Matches))
	}
}

func toString(m map[string]int) string {
	s := ""
	for k := range m {
		s += k + ","
	}

	return s
}
