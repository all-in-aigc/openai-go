package examples

import (
	"time"

	"github.com/all-in-aigc/gpt"
)

var (
	apiKey string
)

func getClient() *gpt.Client {
	cli, _ := gpt.NewClient(&gpt.Options{
		BaseUri: "https://api.openai.com", // you can use a proxy url here
		ApiKey:  apiKey,
		Timeout: 30 * time.Second,
		Debug:   true,
	})

	return cli
}
