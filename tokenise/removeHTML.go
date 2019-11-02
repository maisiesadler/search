package tokenise

import (
	"strings"
)

func removeHTML(ch <-chan string) <-chan string {
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
