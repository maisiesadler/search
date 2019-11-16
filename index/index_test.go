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

func findAndAssertDictionaryResult(t *testing.T, dict Dictionary, searchterm string) *DictionaryResult {
	found, result := dict.Find(searchterm)

	if !found {
		t.Errorf("Could not find word '%v' in dictionary", searchterm)
		t.FailNow()
	}

	return result
}

func findAndAssertOneResult(t *testing.T, idx Index, searchterm string) *Result {
	return findAndAssertSearchTermForOneResult(t, idx, searchterm, searchterm)
}

func findAndAssertSearchTermForOneResult(t *testing.T, idx Index, searchterm string, word string) *Result {
	found, results := idx.Find(searchterm)

	if !found {
		t.Errorf("Could not find word '%v' in index", searchterm)
		t.FailNow()
	}

	if len(results) != 1 {
		t.Errorf("Unexpected number of results in index. Expected 1, actual %v.", len(results))
		t.FailNow()
	}

	result := results[0]

	if result.Word != word {
		t.Errorf("Word returned by index '%v' is not expected word '%v'.", result.Word, word)
	}

	if len(result.Matches) != 1 {
		t.Errorf("Unexpected number of matches (%v) returned by index", result.Matches)
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

func assertOccurencesForValue(t *testing.T, result *DictionaryResult, value string, expectedCount int) {
	if docresult, ok := result.ValueOccurences[value]; ok {
		if docresult != expectedCount {
			t.Error("Unexpected number of matches for doc in Result")
		}
	} else {
		t.Error("Expected docID does not appear in Matches")
	}
}
