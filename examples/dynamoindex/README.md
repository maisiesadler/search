# DynamoIndex

DynamoIndex is an implementation of Index backed by dynamodb.

Credentials from the shared credentials file ~/.aws/credentials
and region from the shared configuration file ~/.aws/config.

## To Add

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

## To Find

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