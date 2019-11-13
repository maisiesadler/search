package preprocess_test

import (
	"testing"

	"github.com/maisiesadler/search/preprocess"
)

func applyPreprocessor(preprocessor preprocess.Preprocessor, unprocessed []string) []string {
	raw := make(chan string)

	go func() {
		defer close(raw)
		for _, token := range unprocessed {
			raw <- token
		}
	}()

	tokens := preprocessor.Process(raw)

	var processed []string
	for token := range tokens {
		processed = append(processed, token)
	}

	return processed
}

func assertArrays(t *testing.T, expected []string, actual []string) {
	if len(actual) != len(expected) {
		t.Errorf("Actual length (%v) is not equal to expected length (%v)", len(actual), len(expected))
		t.FailNow()
	}

	for i, expectedToken := range expected {
		actualToken := actual[i]
		if actualToken != expectedToken {
			t.Errorf("Token at position %v (%v) is not expected value (%v)", i, actualToken, expectedToken)
			t.FailNow()
		}
	}
}
