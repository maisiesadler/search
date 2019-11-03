package tokenise

// Tokeniser processes text into tokens
type Tokeniser interface {
	Tokenise(ch <-chan string) <-chan string
}

// Tokenise takes a channel of string and applies various rules to output tokens
func Tokenise(ch <-chan string) <-chan string {

	processed := RemoveExcluded().Tokenise(ch)
	processed = RemoveHTML().Tokenise(processed)
	processed = ToWord().Tokenise(processed)

	return processed
}

// ToWord removes punctuation and breaks the sentence by whitespace
func ToWord() Tokeniser {
	return createToWordTokeniser()
}

// RemoveExcluded removes an entire block of text if it includes an excluded term
func RemoveExcluded(excludePhrases ...string) Tokeniser {
	return createRemoveExcludedTokeniser(excludePhrases)
}

// RemoveHTML removes any tokens with <>
func RemoveHTML() Tokeniser {
	return createRemoveHTMLTokeniser()
}
