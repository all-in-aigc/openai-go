package gpt

import (
	"time"
)

// Options can set custom options for ChatGPT request client
type Options struct {
	// Debug is used to output debug message
	Debug bool
	// Timeout is used to end http request after timeout duration
	Timeout time.Duration
	// ApiKey is used to authoration
	ApiKey string
	// BaseUri is used to set api baseuri
	BaseUri string
	// Model is used to set default model
	Model string
	// ApiVersion is used to set api version
	ApiVersion string
	// isAzure is a flag for azure openai
	isAzure bool
}
