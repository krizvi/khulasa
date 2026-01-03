package khulasa

import (
	"github.com/jdkato/prose/v2"
	"github.com/sashabaranov/go-openai"
)

// NewTextKhulasa builds a TextKhulasa with the required inputs and optional settings.
func NewTextKhulasa(text string, summaryRate, tgtPerc float64, opts ...Options) (Khulasa, error) {
	k := &TextKhulasa{
		text:             text,
		summaryRate:      summaryRate,
		targetPercentage: tgtPerc,
	}

	for _, opt := range opts {
		if opt != nil {
			opt(k)
		}
	}

	doc, err := prose.NewDocument(text)
	if err != nil {
		return nil, err
	}
	k.doc = doc

	return k, nil
}

// WithText sets the source text for summarization.
func WithText(text string) Options {
	return func(k *TextKhulasa) {
		k.text = text
	}
}

// WithSummaryRate sets the extractive summary rate (0.0-1.0).
func WithSummaryRate(rate float64) Options {
	return func(k *TextKhulasa) {
		k.summaryRate = rate
	}
}

// WithTargetPercentage sets the target percentage of the original size.
func WithTargetPercentage(pct float64) Options {
	return func(k *TextKhulasa) {
		k.targetPercentage = pct
	}
}

// WithKhulasaType sets the summarization method.
func WithKhulasaType(t KhulasaType) Options {
	return func(k *TextKhulasa) {
		k.khulasaType = t
	}
}

// WithExtractive sets the summarization method to Extractive.
func WithExtractive() Options {
	return func(k *TextKhulasa) {
		k.khulasaType = Extractive
	}
}

// WithAbstractiveOpenAI sets the summarization method to AbstractiveOpenAI.
func WithAbstractiveOpenAI() Options {
	return func(k *TextKhulasa) {
		k.khulasaType = AbstractiveOpenAI
	}
}

// WithAbstractiveHuggingFace sets the summarization method to AbstractiveHuggingFace.
func WithAbstractiveHuggingFace() Options {
	return func(k *TextKhulasa) {
		k.khulasaType = AbstractiveHuggingFace
	}
}

// WithHibrid sets the summarization method to Hibrid.
func WithHibrid() Options {
	return func(k *TextKhulasa) {
		k.khulasaType = Hibrid
	}
}

// WithOpenAIKey sets the OpenAI API key.
func WithOpenAIKey(key string) Options {
	return func(k *TextKhulasa) {
		k.openAIKey = key
	}
}

// WithOpenAIBaseURL sets the OpenAI base URL (useful for proxies/compatible APIs).
func WithOpenAIBaseURL(url string) Options {
	return func(k *TextKhulasa) {
		k.openAIBaseURL = url
	}
}

// WithOpenAIModel sets the OpenAI model name.
func WithOpenAIModel(model string) Options {
	return func(k *TextKhulasa) {
		k.openAIModel = model
	}
}

// WithOpenAIClient sets a custom OpenAI client.
func WithOpenAIClient(client *openai.Client) Options {
	return func(k *TextKhulasa) {
		k.openAIClient = client
	}
}

// WithHuggingFaceConfig sets the Hugging Face configuration.
func WithHuggingFaceConfig(cfg HuggingFaceConfig) Options {
	return func(k *TextKhulasa) {
		k.huggingFaceConfig = cfg
	}
}

// GenerateKhulasa implements Khulasa; stubbed for now.
func (k *TextKhulasa) GenerateKhulasa() string {
	return ""
}

// GetResponse implements Khulasa; stubbed for now.
func (k *TextKhulasa) GetResponse() SummaryResponse {
	return SummaryResponse{}
}

// ExtractKeywords implements Khulasa; stubbed for now.
func (k *TextKhulasa) ExtractKeywords(count int) []string {
	return nil
}
