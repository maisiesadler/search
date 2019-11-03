package preprocess

import (
	"testing"
)

func TestLanguageRemovesAccents(t *testing.T) {
	unprocessed := "cliché"
	expected := "cliche"

	processed := preprocessLanguage([]string{unprocessed})

	if len(processed) != 1 {
		t.Error("Processed token array was not expected length")
	}

	if processed[0] != expected {
		t.Errorf("Processed token (%v) was not expected value (%v)", processed[0], expected)
	}
}

func preprocessLanguage(unprocessed []string) []string {
	raw := make(chan string)

	go func() {
		defer close(raw)
		for _, token := range unprocessed {
			raw <- token
		}
	}()

	tokens := language(raw)

	var processed []string
	for token := range tokens {
		processed = append(processed, token)
	}

	return processed
}
