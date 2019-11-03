package preprocess_test

import (
	"search/preprocess"
	"testing"
)

func TestLanguageRemovesAccents(t *testing.T) {
	unprocessed := "cliché"
	expected := "cliche"

	preprocessor := preprocess.Language()
	processed := applyPreprocessor(preprocessor, []string{unprocessed})

	if len(processed) != 1 {
		t.Error("Processed token array was not expected length")
	}

	if processed[0] != expected {
		t.Errorf("Processed token (%v) was not expected value (%v)", processed[0], expected)
	}
}
