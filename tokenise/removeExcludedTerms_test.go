package tokenise_test

import (
	"search/tokenise"
	"testing"
)

func TestRemoveExcludedTerms(t *testing.T) {
	content := []string{
		"Back to main page",
		"this is the content",
		"some footer"}
	expectedtokens := []string{"this is the content"}

	tokeniser := tokenise.RemoveExcluded("Back to main page", "footer")

	words := applyTokeniser(tokeniser, content)

	assertArrays(t, expectedtokens, words)
}
