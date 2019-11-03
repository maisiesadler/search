package preprocess_test

import (
	"search/preprocess"
	"testing"
)

func TestLowercaseReturnsTokenAsLower(t *testing.T) {
	unprocessed := "TOKEN"
	expected := "token"

	preprocessor := preprocess.Lowercase()
	processed := applyPreprocessor(preprocessor, []string{unprocessed})

	if len(processed) != 1 {
		t.Error("Processed token array was not expected length")
	}

	if processed[0] != expected {
		t.Error("Processed token was not expected value")
	}
}
