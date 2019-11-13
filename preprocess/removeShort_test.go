package preprocess_test

import (
	"testing"

	"github.com/maisiesadler/search/preprocess"
)

func TestRemoveShortRemovesShortTokens(t *testing.T) {
	unprocessed := []string{"a", "few", "tokens"}
	expected := []string{"few", "tokens"}

	preprocessor := preprocess.RemoveShort()
	processed := applyPreprocessor(preprocessor, unprocessed)

	assertArrays(t, expected, processed)
}
