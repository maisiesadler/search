package index

// Index is the index
type Index interface {
	Add(docID string, tokens <-chan string)
	AddOne(docID string, token string)
	Find(word string) (bool, []*Result)
	FindOne(word string) (bool, *Result)
	PrintInfo()
}

// Result is the type returned by Find
type Result struct {
	Matches map[string]int
	Word    string
}

// Create returns a kgram Index backed by an in memory dictionary implementation of Index
func Create() Index {
	dictionary := createDictionaryIndex()
	return createKgramIndex(dictionary)
}

// CreateKgramFromIndex returns a kgram Index backed by the provided Index
func CreateKgramFromIndex(index Index) Index {
	return createKgramIndex(index)
}
