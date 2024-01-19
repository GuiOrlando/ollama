package llm

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/jmorganca/ollama/api"
)

type ImageData struct {
	Data []byte `json:"data"`
	ID   int    `json:"id"`
}

var payloadMissing = fmt.Errorf("expected dynamic library payloads not included in this build of ollama")

type prediction struct {
	Content string `json:"content"`
	Model   string `json:"model"`
	Prompt  string `json:"prompt"`
	Stop    bool   `json:"stop"`

	Timings struct {
		PredictedN  int     `json:"predicted_n"`
		PredictedMS float64 `json:"predicted_ms"`
		PromptN     int     `json:"prompt_n"`
		PromptMS    float64 `json:"prompt_ms"`
	}
}

const maxRetries = 3

type PredictOpts struct {
	Prompt  string
	Format  string
	Images  []api.ImageData
	Options api.Options
}

type PredictResult struct {
	Content            string
	Done               bool
	PromptEvalCount    int
	PromptEvalDuration time.Duration
	EvalCount          int
	EvalDuration       time.Duration
}

type TokenizeRequest struct {
	Content string `json:"content"`
}

type TokenizeResponse struct {
	Tokens []int `json:"tokens"`
}

type DetokenizeRequest struct {
	Tokens []int `json:"tokens"`
}

type DetokenizeResponse struct {
	Content string `json:"content"`
}

type EmbeddingRequest struct {
	Content string `json:"content"`
}

type EmbeddingResponse struct {
	Embedding []float64 `json:"embedding"`
}
