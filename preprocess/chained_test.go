package preprocess_test

import (
	"testing"

	"github.com/maisiesadler/search/preprocess"
)

func TestChainedAppliesFunctionsInOrder(t *testing.T) {
	// Arrange
	unprocessed := "test"
	expected12 := "testonetwo"
	expected21 := "testonetwo"

	preprocessor1 := &appendPreprocessor{"one"}
	preprocessor2 := &appendPreprocessor{"two"}

	// Act
	preprocessor12 := preprocess.Chained(preprocessor1, preprocessor2)
	processed12 := applyPreprocessor(preprocessor12, []string{unprocessed})

	preprocessor21 := preprocess.Chained(preprocessor1, preprocessor2)
	processed21 := applyPreprocessor(preprocessor21, []string{unprocessed})

	// Assert
	assertArrays(t, []string{expected12}, processed12)
	assertArrays(t, []string{expected21}, processed21)
}

type appendPreprocessor struct {
	appendword string
}

func (p *appendPreprocessor) Process(ch <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for val := range ch {
			out <- val + p.appendword
		}

		close(out)
	}()

	return out
}
