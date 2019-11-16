package index

import "testing"

func TestKGram_ExactMatch_CanFind(t *testing.T) {
	docID := "123"
	word := "hello"

	idx := createKgramIndexWithOneWord(docID, word)

	result := findAndAssertOneResult(t, idx, word)

	assertCountForDocID(t, result, docID, 1)
}

func TestKGram_DoesNotMatch_CantFind(t *testing.T) {
	docID := "123"
	word := "hello"
	searchterm := "ello"

	idx := createKgramIndexWithOneWord(docID, word)

	found, _ := idx.Find(searchterm)

	if found {
		t.Error("Did not expect to find word in index")
	}
}

func TestKGram_Wildcard_CanFind(t *testing.T) {
	docID := "123"
	word := "hello"
	searchterm := "h*lo"

	idx := createKgramIndexWithOneWord(docID, word)

	result := findAndAssertSearchTermForOneResult(t, idx, searchterm, word)

	assertCountForDocID(t, result, docID, 1)
}

func createKgramIndexWithOneWord(docID string, word string) Index {
	idx := Create()

	tokens := make(chan string)
	go func() {
		defer close(tokens)
		tokens <- word
	}()

	idx.Add(docID, tokens)

	return idx
}
