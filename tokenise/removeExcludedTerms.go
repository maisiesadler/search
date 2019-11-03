package tokenise

import (
	"regexp"
	"strings"
)

type removeExcludedTokeniser struct {
	excludeRegex *regexp.Regexp
}

func createRemoveExcludedTokeniser(excludePhrases []string) Tokeniser {
	excludeRegex := regexp.MustCompile(strings.Join(excludePhrases, "|"))
	return &removeExcludedTokeniser{excludeRegex}
}

func (t *removeExcludedTokeniser) Tokenise(ch <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for val := range ch {
			if !t.excludeRegex.MatchString(val) {
				out <- val
			}
		}

		close(out)
	}()

	return out
}
