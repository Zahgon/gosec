package autofix

const (
	modelAtlasDefault         = "deepseek-ai/deepseek-v4-flash"
	modelAtlasDeepSeekV4Flash = "deepseek-ai/deepseek-v4-flash"
	modelAtlasQwenCoderNext   = "qwen/qwen3-coder-next"
	modelAtlasKimiK26         = "moonshotai/kimi-k2.6"

	defaultAtlasBaseURL = "https://api.atlascloud.ai/v1"
)

type atlasConfig struct {
	Model       string
	APIKey      string `json:"-"`
	BaseURL     string
	MaxTokens   int
	Temperature float64
	SkipSSL     bool
}

func newAtlasClient(config atlasConfig) (GenAIClient, error) {
	_ = "STUB: not implemented"
	return *new(GenAIClient), nil
}

func parseAtlasModel(model string) string { _ = "STUB: not implemented"; return "" }
