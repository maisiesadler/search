package tokenise

func removeEnd(ch <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for val := range ch {
			if val == "Navigation menu" {
				close(out)
				return
			}

			out <- val
		}

		close(out)
	}()

	return out
}
