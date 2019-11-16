package index

import (
	"fmt"
	"strconv"
	"strings"
)

type instance struct {
	token string
	count int
	docID string
}

type kgramIndex struct {
	dictionary Index
	kgrams     map[string]map[string]bool
}

func createKgramIndex(index Index) *kgramIndex {
	return &kgramIndex{kgrams: make(map[string]map[string]bool), dictionary: index}
}

func (ki *kgramIndex) Add(docID string, tokens <-chan string) {
	for token := range tokens {
		ki.AddOne(docID, token)
	}
}

func (ki *kgramIndex) AddOne(docID string, token string) {
	ki.dictionary.AddOne(docID, token)
	s := "$" + token + "$" + token[:1]
	for kgram := range createKgrams(s) {
		if _, ok := ki.kgrams[kgram]; !ok {
			ki.kgrams[kgram] = make(map[string]bool)
		}
		ki.kgrams[kgram][token] = true
	}
}

func createKgrams(s string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 3; i <= len(s); i++ {
			ch <- s[i-3 : i]
		}
		close(ch)
	}()

	return ch
}

func (ki *kgramIndex) Find(word string) (bool, []*Result) {
	var results []map[string]bool

	ok, rotatedWildcard := createWildcard(word)
	if !ok {
		return false, nil
	}

	for kgram := range createKgrams(rotatedWildcard) {
		if tokenmap, ok := ki.kgrams[kgram]; ok {
			results = append(results, tokenmap)
		} else {
			// if any kgrams for word don't exist, does not match
			return false, nil
		}
	}

	if len(results) == 0 {
		return false, nil
	}

	res := booleanAnd(results)
	if len(res) == 0 {
		return false, nil
	}

	var finalRes []*Result

	for _, r := range res {
		if ok, i := ki.dictionary.FindOne(r); ok {
			// i.Word = r
			finalRes = append(finalRes, i)
		}
	}

	return true, finalRes
}

func (ki *kgramIndex) FindOne(word string) (bool, *Result) {

	found, results := ki.Find(word)
	if found && len(results) > 0 {
		return true, results[0]
	}

	return false, nil
}

func createWildcard(s string) (bool, string) {
	wildcardCount := strings.Count(s, "*")
	s = "$" + s
	if wildcardCount == 0 {
		return true, s + "$"
	}
	if wildcardCount == 1 {
		lens := len(s)
		for i := 0; i < lens; i++ {
			s = s[lens-1:] + s[:lens-1]
			if s[lens-1:] == "*" {
				return true, s[:lens-1]
			}
		}
	}

	return false, ""
}

func booleanAnd(results []map[string]bool) []string {

	if len(results) == 0 {
		return []string{}
	}

	// pick result with smallest posting list
	smallestSet := results[0]
	smallestSetLength := len(results[0])
	for _, result := range results[1:] {
		setLength := len(result)
		if setLength < smallestSetLength {
			smallestSet = result
			smallestSetLength = setLength
		}
	}

	// take smallest as first set
	m := smallestSet

	for _, ti := range results {
		for k := range m {
			if _, ok := ti[k]; !ok {
				delete(m, k)
			}
		}
	}

	var s []string
	for k := range m {
		s = append(s, k)
	}

	return s
}

func (ki *kgramIndex) PrintInfo() {
	fmt.Println("added " + strconv.Itoa(len(ki.kgrams)) + " kgrams")
}
