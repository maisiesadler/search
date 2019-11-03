package preprocess

import "testing"

func TestLowercaseReturnsTokenAsLower(t *testing.T) {
	unprocessed := "TOKEN"
	expected := "token"

	processed := preprocessLowercase([]string{unprocessed})

	if len(processed) != 1 {
		t.Error("Processed token array was not expected length")
	}

	if processed[0] != expected {
		t.Error("Processed token was not expected value")
	}
}

func preprocessLowercase(unprocessed []string) []string {
	raw := make(chan string)

	go func() {
		defer close(raw)
		for _, token := range unprocessed {
			raw <- token
		}
	}()

	tokens := lowercase(raw)

	var processed []string
	for token := range tokens {
		processed = append(processed, token)
	}

	return processed
}
