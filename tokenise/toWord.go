package tokenise

import (
	"regexp"
	"strings"
)

var charsToRemoveReg, _ = regexp.Compile("\\[\\d+\\]|[();:.,'\"^\\[\\]]")
var charsToWhitespaceReg, _ = regexp.Compile("[-]")
var onlyNumberReg, _ = regexp.Compile("^\\d+\\.?\\d*$")

func toWords(ch <-chan string) <-chan string {
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
