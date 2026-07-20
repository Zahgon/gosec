package autofix

import (
	"context"

	"github.com/anthropics/anthropic-sdk-go"
)

const (
	ModelClaudeOpus4_7   = anthropic.ModelClaudeOpus4_7
	ModelClaudeOpus4_6   = anthropic.ModelClaudeOpus4_6
	ModelClaudeSonnet4_6 = anthropic.ModelClaudeSonnet4_6
	ModelClaudeOpus4_5   = anthropic.ModelClaudeOpus4_5_20251101
	ModelClaudeSonnet4_5 = anthropic.ModelClaudeSonnet4_5_20250929
	ModelClaudeHaiku4_5  = anthropic.ModelClaudeHaiku4_5_20251001
)

var _ GenAIClient = (*claudeWrapper)(nil)

type claudeWrapper struct {
	client anthropic.Client
	model  anthropic.Model
}

func NewClaudeClient(model, apiKey string) (GenAIClient, error) {
	_ = "STUB: not implemented"
	return *new(GenAIClient), nil
}

func (c *claudeWrapper) GenerateSolution(ctx context.Context, prompt string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func parseAnthropicModel(model string) anthropic.Model {
	_ = "STUB: not implemented"
	return *new(anthropic.Model)
}
