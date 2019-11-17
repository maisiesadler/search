package index

import "testing"

func TestCanAddToDictionaryIndex(t *testing.T) {
	dictionary := createDictionaryIndex()
	idx := createIndexFromDictionary(dictionary)

	docID := "doc-id"
	token := "word"

	testAddToIdx(idx, docID, token)

	found, results := idx.Find(token)

	result := assertAndGetOneResult(t, found, results)

	if len(result.Matches) != 1 {
		t.Errorf("Did not find expected number of matches. Expected 1, actual %v.", len(result.Matches))
		t.FailNow()
	}

	if _, ok := result.Matches[docID]; !ok {
		t.Error("Did not find doc-id in matches.")
		t.FailNow()
	}
}

func TestCanAddMultipleToDictionaryIndex(t *testing.T) {
	dictionary := createDictionaryIndex()
	idx := createIndexFromDictionary(dictionary)

	docID := "doc-id"
	docID2 := "doc-id-2"
	token := "word"

	testAddToIdx(idx, docID, token)
	testAddToIdx(idx, docID2, token)

	found, results := idx.Find(token)

	result := assertAndGetOneResult(t, found, results)

	if len(result.Matches) != 2 {
		t.Errorf("Did not find expected number of matches. Expected 2, actual %v.", len(result.Matches))
		t.FailNow()
	}

	if _, ok := result.Matches[docID]; !ok {
		t.Error("Did not find doc-id in matches.")
		t.FailNow()
	}

	if _, ok := result.Matches[docID2]; !ok {
		t.Error("Did not find doc-id-2 in matches.")
		t.FailNow()
	}
}

func assertAndGetOneResult(t *testing.T, found bool, results []*Result) *Result {
	if !found {
		t.Error("Did not find any results")
		t.FailNow()
	}

	if len(results) != 1 {
		t.Errorf("Did not find expected number of results. Expected 1, actual %v.", len(results))
		t.FailNow()
	}

	return results[0]
}

func testAddToIdx(idx Index, docID string, token string) {
	tokens := make(chan string)
	go func() {
		defer close(tokens)
		tokens <- token
	}()
	idx.Add(docID, tokens)
}
