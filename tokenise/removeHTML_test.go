package tokenise_test

import (
	"testing"

	"github.com/maisiesadler/search/tokenise"
)

func TestRemoveHTML(t *testing.T) {
	content := []string{
		"<div>",
		"this is the content",
		"</div>"}
	expectedtokens := []string{"this is the content"}

	tokeniser := tokenise.RemoveHTML()

	words := applyTokeniser(tokeniser, content)

	assertArrays(t, expectedtokens, words)
}
