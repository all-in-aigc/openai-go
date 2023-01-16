package examples

import (
	"time"

	"github.com/chatgp/gpt3"
)

var (
	apiKey string
)

func getClient() *gpt3.Client {
	cli, _ := gpt3.NewClient(&gpt3.Options{
		ApiKey:  apiKey,
		Timeout: 30 * time.Second,
		Debug:   true,
	})

	return cli
}
