package tokenise_test

import (
	"search/tokenise"
	"testing"
)

func applyTokeniser(t tokenise.Tokeniser, arr []string) []string {
	raw := make(chan string)

	go func() {
		defer close(raw)
		for _, sentence := range arr {
			raw <- sentence
		}
	}()

	tokenChan := t.Tokenise(raw)

	var tokens []string
	for token := range tokenChan {
		tokens = append(tokens, token)
	}

	return tokens
}

func assertArrays(t *testing.T, expected []string, actual []string) {
	if len(actual) != len(expected) {
		t.Errorf("Actual length (%v) is not equal to expected length (%v)", len(actual), len(expected))
		t.FailNow()
	}

	for i, expectedToken := range expected {
		actualToken := actual[i]
		if actualToken != expectedToken {
			t.Errorf("Token at position %v (%v) is not equal to (%v)", i, actualToken, expectedToken)
			t.FailNow()
		}
	}
}
