package preprocess

type chainedPreprocessor struct {
	preprocessors []Preprocessor
}

func createChainedPreprocessor(preprocessors []Preprocessor) Preprocessor {
	return &chainedPreprocessor{preprocessors}
}

func (p *chainedPreprocessor) Process(ch <-chan string) <-chan string {
	out := ch

	for _, preprocessor := range p.preprocessors {
		out = preprocessor.Process(out)
	}

	return out
}
