package preprocess

// Preprocess takes a channel of string and applies various rules to group tokens
func Preprocess(ch <-chan string) <-chan string {
	processed := Lowercase().Process(ch)
	processed = RemoveShort().Process(processed)
	processed = Language().Process(processed)
	processed = Stemming().Process(processed)

	return processed
}

// Preprocessor will apply token normalisation rules
type Preprocessor interface {
	Process(ch <-chan string) <-chan string
}

// Language removes accents and diacratics
func Language() Preprocessor {
	return createLanguageProcessor()
}

// Lowercase returns token in lowercase
func Lowercase() Preprocessor {
	return createLowercaseProcessor()
}

// RemoveShort excludes tokens which are too short for indexing
func RemoveShort() Preprocessor {
	return createRemoveShortProcessor()
}

// RemoveStopWords removes configurable common tokens, e.g. 'of|and|as'
func RemoveStopWords(stopWords []string) Preprocessor {
	return createRemoveStopWordsProcessor(stopWords)
}

// Stemming attempts to group families of words by replacing suffixes with one common suffix
func Stemming() Preprocessor {
	return createStemmingProcessor()
}
