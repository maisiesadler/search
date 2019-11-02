package tokenise

import "testing"

func TestToWordCreatesATokenForEachWord(t *testing.T) {
	sentence := "this is a sentence about nothing"
	expectedtokens := []string{"this", "is", "a", "sentence", "about", "nothing"}

	words := processArrayToWords([]string{sentence})

	assertArrays(t, expectedtokens, words)
}

func TestDuplicateWordsAreAddedTwice(t *testing.T) {
	sentence := "this is a sentence about nothing"
	expectedtokens := []string{
		"this", "is", "a", "sentence", "about", "nothing",
		"this", "is", "a", "sentence", "about", "nothing"}

	words := processArrayToWords([]string{sentence, sentence})

	assertArrays(t, expectedtokens, words)
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

func processArrayToWords(arr []string) []string {
	raw := make(chan string)

	go func() {
		defer close(raw)
		for _, sentence := range arr {
			raw <- sentence
		}
	}()

	wordChan := toWords(raw)

	var words []string
	for word := range wordChan {
		words = append(words, word)
	}

	return words
}
