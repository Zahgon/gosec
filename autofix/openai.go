package autofix

import (
	"context"

	"github.com/openai/openai-go/v3"
)

const (
	ModelGPT5_4          = openai.ChatModelGPT5_4
	ModelGPT5_4Mini      = openai.ChatModelGPT5_4Mini
	ModelGPT5_4Nano      = openai.ChatModelGPT5_4Nano
	DefaultOpenAIBaseURL = "https://api.openai.com/v1"
)

var _ GenAIClient = (*openaiWrapper)(nil)

type OpenAIConfig struct {
	Model       string
	APIKey      string `json:"-"`
	BaseURL     string
	MaxTokens   int
	Temperature float64
	SkipSSL     bool
}

type openaiWrapper struct {
	client      openai.Client
	model       openai.ChatModel
	maxTokens   int
	temperature float64
}

func NewOpenAIClient(config OpenAIConfig) (GenAIClient, error) {
	_ = "STUB: not implemented"
	return *new(GenAIClient), nil
}

func (o *openaiWrapper) GenerateSolution(ctx context.Context, prompt string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func parseOpenAIModel(model string) openai.ChatModel {
	_ = "STUB: not implemented"
	return *new(openai.ChatModel)
}
