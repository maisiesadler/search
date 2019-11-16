package index

type dictionaryIndex struct {
	tokens map[string]map[string]int
}

func createDictionaryIndex() Dictionary {
	return &dictionaryIndex{tokens: make(map[string]map[string]int)}
}

func (di *dictionaryIndex) Add(key string, value string) {
	if _, ok := di.tokens[key]; !ok {
		di.tokens[key] = make(map[string]int)
	}
	di.tokens[key][value]++
}

func (di *dictionaryIndex) Find(key string) (bool, *DictionaryResult) {
	if docIDs, ok := di.tokens[key]; ok {
		return true, &DictionaryResult{Key: key, ValueOccurences: docIDs}
	}

	return false, nil
}

func keys(m map[string]int) []*string {
	keys := []*string{}
	for k := range m {
		keys = append(keys, &k)
	}

	return keys
}
