package preprocess

import (
	"strings"
)

type lowercaseProcessor struct {
}

func createLowercaseProcessor() Preprocessor {
	return &lowercaseProcessor{}
}

func (l *lowercaseProcessor) Process(ch <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for val := range ch {
			out <- strings.ToLower(val)
		}

		close(out)
	}()

	return out
}
