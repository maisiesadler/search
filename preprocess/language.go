package preprocess

import (
	"regexp"
	"strings"
)

var langRegA, _ = regexp.Compile("àä")
var langRegE, _ = regexp.Compile("ëé")
var langRegO, _ = regexp.Compile("ô")

func language(ch <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for val := range ch {
			val = langRegA.ReplaceAllString(val, "a")
			val = langRegE.ReplaceAllString(val, "e")
			val = langRegO.ReplaceAllString(val, "o")

			out <- strings.ToLower(val)
		}

		close(out)
	}()

	return out
}
