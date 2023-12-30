package examples

import (
	"time"

	"github.com/all-in-aigc/openai-go"
)

var (
	apiBaseUri string
	apiKey     string
	apiVersion string
	model      string
)

func getClient() *openai.Client {
	cli, _ := openai.NewClient(&openai.Options{
		BaseUri:    apiBaseUri,
		ApiKey:     apiKey,
		Timeout:    30 * time.Second,
		Debug:      true,
		ApiVersion: apiVersion,
		Model:      model,
	})

	return cli
}
