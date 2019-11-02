package index

import "testing"

func TestDictionary_ExactMatch_CanFind(t *testing.T) {
	docID := "123"
	word := "hello"

	idx := createDictionaryWithOneWord(docID, word)

	result := findAndAssertOneResult(t, idx, word)

	assertCountForDocID(t, result, docID, 1)
}

func TestDictionary_DifferentId_CantFind(t *testing.T) {
	docID := "123"
	word := "hello"
	anotherword := "nothello"

	idx := createDictionaryWithOneWord(docID, word)

	found, _ := idx.Find(anotherword)

	if found {
		t.Error("Did not expect to find word in index")
	}
}

func createDictionaryWithOneWord(docID string, word string) *dictionaryIndex {
	idx := createDictionaryIndex()

	tokens := make(chan string)
	defer close(tokens)

	go idx.Add(docID, tokens)

	tokens <- word

	return idx
}
