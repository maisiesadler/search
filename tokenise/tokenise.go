package tokenise

// Tokenise takes a channel of string and applies various rules to output tokens
func Tokenise(ch <-chan string) <-chan string {

	processed := removeExcludedTerms(ch)
	processed = removeHTML(processed)
	processed = toWords(processed)

	return processed
}
