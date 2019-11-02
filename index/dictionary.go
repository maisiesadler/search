package index

import (
	"fmt"
	"strconv"
)

type dictionaryIndex struct {
	tokens map[string]map[string]int
}

func createDictionaryIndex() *dictionaryIndex {
	return &dictionaryIndex{tokens: make(map[string]map[string]int)}
}

func (di *dictionaryIndex) Add(docID string, tokens <-chan string) {
	for token := range tokens {
		di.addone(docID, token)
	}
}

func (di *dictionaryIndex) addone(docID string, token string) {
	if docIDs, ok := di.tokens[token]; !ok {
		docIDs = make(map[string]int)
		di.tokens[token] = docIDs
	}
	di.tokens[token][docID]++
}

func (di *dictionaryIndex) Find(word string) (bool, []*Result) {
	if ok, result := di.FindOne(word); ok {
		return true, []*Result{result}
	}

	return false, nil
}

func (di *dictionaryIndex) FindOne(word string) (bool, *Result) {
	if docIDs, ok := di.tokens[word]; ok {
		return true, &Result{Word: word, Matches: docIDs}
	}

	return false, nil
}

func (di *dictionaryIndex) PrintInfo() {
	fmt.Println("Keys in index:" + strconv.Itoa(len(di.tokens)))
}
