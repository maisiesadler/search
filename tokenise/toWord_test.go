package tokenise_test

import (
	"testing"

	"github.com/maisiesadler/search/tokenise"
)

func TestToWordCreatesATokenForEachWord(t *testing.T) {
	sentence := "this is a sentence about nothing"
	expectedtokens := []string{"this", "is", "a", "sentence", "about", "nothing"}

	tokeniser := tokenise.ToWord()

	words := applyTokeniser(tokeniser, []string{sentence})

	assertArrays(t, expectedtokens, words)
}

func TestDuplicateWordsAreAddedTwice(t *testing.T) {
	sentence := "this is a sentence about nothing"
	expectedtokens := []string{
		"this", "is", "a", "sentence", "about", "nothing",
		"this", "is", "a", "sentence", "about", "nothing"}

	tokeniser := tokenise.ToWord()

	words := applyTokeniser(tokeniser, []string{sentence, sentence})

	assertArrays(t, expectedtokens, words)
}
