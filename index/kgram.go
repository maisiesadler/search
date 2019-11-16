package index

import (
	"strings"
)

type instance struct {
	token string
	count int
	docID string
}

type kgramIndex struct {
	dictionary Dictionary
}

func (ki *kgramIndex) addTokenToDocumentID(token string, docID string) {
	ki.dictionary.Add("-"+token, docID)
}

func (ki *kgramIndex) findDocumentIDsFromToken(token string) (bool, *DictionaryResult) {
	return ki.dictionary.Find("-" + token)
}

func (ki *kgramIndex) addKgram(kgram string, token string) {
	ki.dictionary.Add("_"+kgram, token)
}

func (ki *kgramIndex) findTokensFromKgram(kgram string) (bool, *DictionaryResult) {
	return ki.dictionary.Find("_" + kgram)
}

func createKgramIndex(dictionary Dictionary) *kgramIndex {
	return &kgramIndex{dictionary: dictionary}
}

func (ki *kgramIndex) Add(docID string, tokens <-chan string) {
	for token := range tokens {
		ki.addOne(docID, token)
	}
}

func (ki *kgramIndex) addOne(docID string, token string) {
	ki.dictionary.Add(docID, token)
	s := "$" + token + "$" + token[:1]
	for kgram := range createKgrams(s) {
		ki.addKgram(kgram, token)
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
	var results []map[string]int

	ok, rotatedWildcard := createWildcard(word)
	if !ok {
		return false, nil
	}

	for kgram := range createKgrams(rotatedWildcard) {
		if ok, tokens := ki.findTokensFromKgram(kgram); ok {
			results = append(results, tokens.ValueOccurences)
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
		if ok, occurences := ki.dictionary.Find(r); ok {
			// i.Word = r
			result := &Result{
				Word:    r,
				Matches: occurences.ValueOccurences,
			}
			finalRes = append(finalRes, result)
		}
	}

	return true, finalRes
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

func booleanAnd(results []map[string]int) []string {

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
