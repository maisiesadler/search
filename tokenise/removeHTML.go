package tokenise

import (
	"strings"
)

type removeHTMLTokeniser struct {
}

func createRemoveHTMLTokeniser() Tokeniser {
	return &removeHTMLTokeniser{}
}

func (t *removeHTMLTokeniser) Tokenise(ch <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for val := range ch {
			if strings.Contains(val, "<") {
				continue
			}

			out <- strings.ToLower(val)
		}

		close(out)
	}()

	return out
}
