package examples

import (
	"fmt"
	"log"
)

func ExampleCreateCompletion() {
	cli := getClient()

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

	// Output: xxx
}
