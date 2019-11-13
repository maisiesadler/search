package preprocess_test

import (
	"testing"

	"github.com/maisiesadler/search/preprocess"
)

func TestRemoveStopWordsRemovesConfigured(t *testing.T) {
	stopWords := []string{"and", "of"}
	unprocessed := []string{"pizza", "and", "chips"}
	expected := []string{"pizza", "chips"}

	preprocessor := preprocess.RemoveStopWords(stopWords)
	processed := applyPreprocessor(preprocessor, unprocessed)

	assertArrays(t, expected, processed)
}

func TestEmptyStopWordsDoesNotThrowException(t *testing.T) {
	stopWords := []string{}
	unprocessed := []string{"pizza", "and", "chips"}
	// regexp is an empty string, so matches everything
	expected := []string{}

	preprocessor := preprocess.RemoveStopWords(stopWords)
	processed := applyPreprocessor(preprocessor, unprocessed)

	assertArrays(t, expected, processed)
}
