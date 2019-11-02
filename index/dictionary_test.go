package index

import "testing"

func TestDictionary_ExactMatch_CanFind(t *testing.T) {
	docID := "123"
	word := "hello"

	idx := createDictionaryIndexWithOneWord(docID, word)

	result := findAndAssertOneResult(t, idx, word)

	assertCountForDocID(t, result, docID, 1)
}

func TestDictionary_DoesNotMatch_CantFind(t *testing.T) {
	docID := "123"
	word := "hello"
	searchterm := "nothello"

	idx := createDictionaryIndexWithOneWord(docID, word)

	found, _ := idx.Find(searchterm)

	if found {
		t.Error("Did not expect to find word in index")
	}
}

func createDictionaryIndexWithOneWord(docID string, word string) *dictionaryIndex {
	idx := createDictionaryIndex()

	tokens := make(chan string)
	defer close(tokens)

	go idx.Add(docID, tokens)

	tokens <- word

	return idx
}
