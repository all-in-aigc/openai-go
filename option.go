package gpt3

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
}
