package preprocess

import (
	"strings"
)

func lowercase(ch <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for val := range ch {
			out <- strings.ToLower(val)
		}

		close(out)
	}()

	return out
}
