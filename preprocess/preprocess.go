package preprocess

// Preprocess takes a channel of string and applies various rules to group tokens
func Preprocess(ch <-chan string) <-chan string {
	processed := lowercase(ch)
	processed = removeShort(processed)
	processed = language(processed)
	processed = stemming(processed)

	return processed
}
