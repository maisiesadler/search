package index

// Index is the index
type Index interface {
	Add(docID string, tokens <-chan string)
	Find(word string) (bool, []*Result)
	PrintInfo()
}

// Result is the type returned by Find
type Result struct {
	Matches map[string]int
	Word    string
}

// Create returns the current implementation of Index
func Create() Index {
	return createKgramIndex()
}
