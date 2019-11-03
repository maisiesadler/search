package preprocess

import (
	"regexp"
)

// crude implementation to see how getting these words to go to one version affects the performance/output
// operate operating operates operation operative operatives operational

var stemmingReg, _ = regexp.Compile("(\\w+)(ate|ating|ates|ation|ative|atives|ational)\\b")

type stemmingProcessor struct {
}

func createStemmingProcessor() Preprocessor {
	return &stemmingProcessor{}
}

func (p *stemmingProcessor) Process(ch <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for val := range ch {
			if stemmingReg.MatchString(val) {
				t := stemmingReg.FindStringSubmatch(val)
				val = t[1] + "ate"
			}
			out <- val
		}

		close(out)
	}()

	return out
}
