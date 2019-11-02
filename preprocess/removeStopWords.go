package preprocess

import "regexp"

var stopWordsReg, _ = regexp.Compile("of|and|as")

func removeStopWords(ch <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for val := range ch {
			if !stopWordsReg.MatchString(val) {
				out <- val
			}
		}

		close(out)
	}()

	return out
}
