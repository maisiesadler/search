package index

type dictionary struct {
	tokens map[string]map[string]int
}

func createDictionaryIndex() Dictionary {
	return &dictionary{tokens: make(map[string]map[string]int)}
}

func (di *dictionary) Add(key string, value string) {
	if values, ok := di.tokens[key]; !ok {
		values = make(map[string]int)
		di.tokens[key] = values
	}
	di.tokens[key][value]++
}

func (di *dictionary) Find(key string) (bool, *DictionaryResult) {
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
