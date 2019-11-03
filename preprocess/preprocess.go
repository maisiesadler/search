package preprocess

// Preprocess takes a channel of string and applies various rules to group tokens
func Preprocess(ch <-chan string) <-chan string {
	chained := Chained(Lowercase(), RemoveShort(), Language(), Stemming())
	return chained.Process(ch)
}

// Preprocessor will apply token normalisation rules
type Preprocessor interface {
	Process(ch <-chan string) <-chan string
}

// Chained creates a preprocessor that applies each processor one by one
func Chained(preprocessors ...Preprocessor) Preprocessor {
	return createChainedPreprocessor(preprocessors)
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
