package preprocess

import (
	"regexp"
	"strings"
)

type removeStopWordsProcessor struct {
	stopWords *regexp.Regexp
}

func createRemoveStopWordsProcessor(stopwords []string) Preprocessor {
	stopWordsReg := regexp.MustCompile(strings.Join(stopwords, "|"))
	return &removeStopWordsProcessor{stopWordsReg}
}

func (p *removeStopWordsProcessor) Process(ch <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for val := range ch {
			if !p.stopWords.MatchString(val) {
				out <- val
			}
		}

		close(out)
	}()

	return out
}
