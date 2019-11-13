package tokenise

import (
	"regexp"
	"strings"
)

var numbersReg, _ = regexp.Compile("\\[\\d+\\]")

type removeNumbersTokeniser struct {
}

func createRemoveNumbersTokeniser() Tokeniser {
	return &removeNumbersTokeniser{}
}

func (t *removeNumbersTokeniser) Tokenise(ch <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for val := range ch {
			if onlyNumberReg.MatchString(val) {
				continue
			}
			val = charsToRemoveReg.ReplaceAllString(val, "")
			val = charsToWhitespaceReg.ReplaceAllString(val, " ")
			for _, s := range strings.Split(val, " ") {
				if s != "" {
					out <- s
				}
			}
		}

		close(out)
	}()

	return out
}
