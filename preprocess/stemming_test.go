package preprocess_test

import (
	"testing"

	"github.com/maisiesadler/search/preprocess"
)

func TestStemmingMatchesExpectedWords(t *testing.T) {
	unprocessed := []string{"operate", "operating", "operates", "operation", "operative", "operatives", "operational"}
	expected := []string{"operate", "operate", "operate", "operate", "operate", "operate", "operate"}

	preprocessor := preprocess.Stemming()
	processed := applyPreprocessor(preprocessor, unprocessed)

	assertArrays(t, expected, processed)
}
