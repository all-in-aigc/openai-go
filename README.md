# gpt3

openai gpt-3 sdk, written by golang.

## Preparation

login the [openai official website](https://beta.openai.com/account/api-keys), and get your own API_KEY.

![](./images/get_api_key.jpg)

## Quick Start

1. install gpt3 sdk

```shell
go get -u github.com/chatgp/gpt3
```

2. request api with gpt-3 client

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/chatgp/gpt3"
)

func main() {
	apiKey := "xxx"

	// new gpt-3 client
	cli, _ := gpt3.NewClient(&gpt3.Options{
		ApiKey:  apiKey,
		Timeout: 30 * time.Second,
		Debug:   true,
	})

	// request api
	uri := "/v1/models"

	res, err := cli.Get(uri)

	if err != nil {
		log.Fatalf("request api failed: %v", err)
	}

	for _, v := range res.Get("data").Array() {
		fmt.Printf("model id: %s\n", v.Get("id").String())
	}
}
```

> see available apis in [OpenAI documents](https://beta.openai.com/docs/api-reference/completions/create)

## Examples

there are some examples under the `examples` folder, check and see how to request other apis.

- [Create chat completion](https://platform.openai.com/docs/api-reference/chat/create)

```go
cli := getClient()

uri := "/v1/chat/completions"
params := map[string]interface{}{
	"model":       "gpt-3.5-turbo",
	"messages":      []map[string]interface{}{
		{"role": "user", "content": "hello"},
	},
}

res, err := cli.Post(uri, params)
if err != nil {
	log.Fatalf("request api failed: %v", err)
}

message := res.Get("choices.0.message.content").String()

fmt.Printf("message is: %s", message)
// Output: xxx
```

- [Create Completion](https://beta.openai.com/docs/api-reference/completions/create)

```go
uri := "/v1/completions"
params := map[string]interface{}{
	"model":       "text-davinci-003",
	"prompt":      "say hello three times",
	"max_tokens":  2048,
	"temperature": 0.9,
	"n":           1,
	"stream":      false,
}

res, err := cli.Post(uri, params)

if err != nil {
	log.Fatalf("request api failed: %v", err)
}

fmt.Println(res.GetString("choices.0.text"))
```

- [Create Edit](https://beta.openai.com/docs/api-reference/edits/create)

```go
uri := "/v1/edits"
params := map[string]interface{}{
	"model":       "text-davinci-edit-001",
	"input":       "Are you hapy today?",
	"instruction": "fix mistake",
	"temperature": 0.9,
	"n":           1,
}

res, err := cli.Post(uri, params)

if err != nil {
	log.Fatalf("request api failed: %v", err)
}

fmt.Println(res.GetString("choices.0.text"))
```

- [Create Image](https://beta.openai.com/docs/api-reference/images/create)

```go
uri := "/v1/images/generations"
params := map[string]interface{}{
	"prompt":          "a beautiful girl with big eyes",
	"n":               1,
	"size":            "256x256",
	"response_format": "url",
}

res, err := cli.Post(uri, params)

if err != nil {
	log.Fatalf("request api failed: %v", err)
}

fmt.Println(res.GetString("data.0.url"))
```

## Communication

- Telegram Group: [ChatGPT Creators](https://t.me/+YkEGeRxB5Q0zODY1)

- Discord Server: [ChatGPT Creators](https://discord.gg/qWshJnJs)

- Wechat Group: [ChatGPT Creators](https://work.weixin.qq.com/gm/66944e9bd30628e9270c980bc756663d)

![](./images/wechat_group.jpg)
