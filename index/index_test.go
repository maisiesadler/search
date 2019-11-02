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

	if docresult, ok := result.Matches[docID]; ok {
		if docresult != 1 {
			t.Error("Unexpected number of matches for doc in Result")
		}
	} else {
		t.Error("Expected docID does not appear in Matches")
	}
}
