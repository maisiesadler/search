package preprocess

func removeShort(ch <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for val := range ch {
			if len(val) > 1 {
				out <- val
			}
		}

		close(out)
	}()

	return out
}
