package index

import "testing"

func TestAddToIndex(t *testing.T) {
	idx := Create()

	tokens := make(chan string)
	defer close(tokens)

	docID := "123"
	word := "hello"

	go idx.Add(docID, tokens)

	tokens <- word

	result := findAndAssertOneResult(t, idx, word)

	assertCountForDocID(t, result, docID, 1)
}

func findAndAssertOneResult(t *testing.T, idx Index, searchterm string) *Result {
	return findAndAssertSearchTermForOneResult(t, idx, searchterm, searchterm)
}

func findAndAssertSearchTermForOneResult(t *testing.T, idx Index, searchterm string, word string) *Result {
	found, results := idx.Find(word)

	if !found {
		t.Error("Could not find word in index")
	}

	if len(results) != 1 {
		t.Error("Unexpected number of results in index")
		t.FailNow()
	}

	result := results[0]

	if result.Word != word {
		t.Error("Word returned by index is not expected word")
	}

	if len(result.Matches) != 1 {
		t.Error("Unexpected number of matches returned by index")
	}

	return result
}

func assertCountForDocID(t *testing.T, result *Result, docID string, expectedCount int) {
	if docresult, ok := result.Matches[docID]; ok {
		if docresult != expectedCount {
			t.Error("Unexpected number of matches for doc in Result")
		}
	} else {
		t.Error("Expected docID does not appear in Matches")
	}
}
