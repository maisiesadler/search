package tokenise

import (
	"regexp"
)

var r, _ = regexp.Compile("Jump to navigation|From Wikipedia, the free encyclopedia|Jump to search")

func removeExcludedTerms(ch <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for val := range ch {
			if !r.MatchString(val) {
				out <- val
			}
		}

		close(out)
	}()

	return out
}
