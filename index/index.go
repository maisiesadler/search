package index

// Index is the index
type Index interface {
	Add(docID string, tokens <-chan string)
	Find(word string) (bool, []*Result)
}

// Dictionary is an index with a 1-1 mapping between key and value
type Dictionary interface {
	Add(key string, value string)
	Find(key string) (bool, *DictionaryResult)
}

// Result is the type returned by Find
type Result struct {
	Matches map[string]int // docID to match value
	Word    string
}

// DictionaryResult is the type returned by Dictionary Find
type DictionaryResult struct {
	ValueOccurences map[string]int // docID to match value
	Key             string
}

// Create returns a kgram Index backed by an in memory implementation of Dictionary
func Create() Index {
	dictionary := createDictionaryIndex()
	return createKgramIndex(dictionary)
}

// CreateKgramFromIndex returns a kgram Index backed by the provided Dictionary
func CreateKgramFromIndex(dictionary Dictionary) Index {
	return createKgramIndex(dictionary)
}

// CreateIndexFromDictionary returns an Index backed by the provided Dictionary with no extra functionality
func CreateIndexFromDictionary(dictionary Dictionary) Index {
	return createIndexFromDictionary(dictionary)
}
