# Search

Search is 3 packages that can be used to index and retreive documents by ID.

Each document is tokenised, preprocessed and added to the index by ID.
The index can then be searched to retreive documents matching by ID.

## [Tokenise](tokenise)

Takes a channel of strings and converts each string into tokens to be used in the preprocessor.

## [Preprocess](preprocess)

Preprocess takes each token and applies normalization rules to group tokens with superficial differences.

## [Index](index)

The normalized tokens can then be added to the index, with a reference to the documentID.
When the index is queried by a term it returns a list of documentIDs containing that term.

## Usage

### To Add

```
idx, _ := CreateDynamoIndex()
tokens := make(chan string)
go func() {
    defer close(tokens)
    tokens <- "here"
    tokens <- "are"
    tokens <- "some"
    tokens <- "tokens"
}()
idx.Add("doc-id", tokens)
```

### To Find

```
found, results := idx.Find("are")
if found {
    fmt.Printf("Found %v result\n", len(results))
} else {
    fmt.Printf("Did not find any results\n")
}
for _, result := range results {
    fmt.Printf("- '%v' = %v\n", result.Word, resultsToString(result.Matches))
}
```

Output: 
Found 1 result
- 'are' = doc-id,

Where resultsToString is:
```
func resultsToString(m map[string]int) string {
	s := ""
	for k := range m {
		s += k + ","
	}

	return s
}
```