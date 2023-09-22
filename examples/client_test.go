package examples

import (
	"time"

	"github.com/all-in-aigc/gpt"
)

var (
	apiBaseUri string
	apiKey     string
	apiVersion string
	model      string
)

func getClient() *gpt.Client {
	cli, _ := gpt.NewClient(&gpt.Options{
		BaseUri:    apiBaseUri,
		ApiKey:     apiKey,
		Timeout:    30 * time.Second,
		Debug:      true,
		ApiVersion: apiVersion,
		Model:      model,
	})

	return cli
}
