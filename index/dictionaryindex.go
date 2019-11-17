package index

type indexUsingDictionary struct {
	dictionary Dictionary
}

func createIndexFromDictionary(dictionary Dictionary) Index {
	return &indexUsingDictionary{dictionary}
}

func (idx *indexUsingDictionary) Add(docID string, tokens <-chan string) {
	for token := range tokens {
		idx.dictionary.Add(token, docID)
	}
}
func (idx *indexUsingDictionary) Find(word string) (bool, []*Result) {
	found, dictResult := idx.dictionary.Find(word)
	if found {
		result := &Result{
			Word:    dictResult.Key,
			Matches: dictResult.ValueOccurences,
		}
		return true, []*Result{result}
	}

	return false, []*Result{}
}
