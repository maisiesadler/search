package preprocess

type removeShortProcessor struct {
}

func createRemoveShortProcessor() Preprocessor {
	return &removeShortProcessor{}
}

func (p *removeShortProcessor) Process(ch <-chan string) <-chan string {
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
