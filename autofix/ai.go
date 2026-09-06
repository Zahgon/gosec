package autofix

import (
	"context"
	"time"

	"github.com/securego/gosec/v2/issue"
)

const (
	AIProviderFlagHelp = `AI API provider to generate auto fixes to issues. Valid options are:
	- atlas (Atlas Cloud default), atlas-deepseek-v4-flash, atlas-qwen3-coder-next, atlas-kimi-k2.6, atlas:<model-id>;
- gemini-3-pro-preview (gemini, default), gemini-2.5-pro, gemini-2.5-flash, gemini-2.5-flash-lite;
- claude-sonnet-4-6 (claude, default), claude-opus-4-7, claude-opus-4-6, claude-sonnet-4-5, claude-opus-4-5, claude-haiku-4-5;
- gpt-5.4 (openai, default), gpt-5.4-mini, gpt-5.4-nano`

	AIPrompt = `Provide a brief explanation and a solution to fix this security issue
  in Go programming language: %q.
  Answer in markdown format and keep the response limited to 200 words.`

	timeout = 30 * time.Second
)

type GenAIClient interface {
	GenerateSolution(ctx context.Context, prompt string) (string, error)
}

func GenerateSolution(model, aiAPIKey, baseURL string, skipSSL bool, issues []*issue.Issue) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func generateSolution(client GenAIClient, issues []*issue.Issue) error {
	_ = "STUB: not implemented"
	return nil
}
