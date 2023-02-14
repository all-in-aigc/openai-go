package examples

import (
	"fmt"
	"log"
	"strings"

	"github.com/tidwall/gjson"
)

func ExampleCreateCompletion() {
	cli := getClient()

	uri := "/v1/completions"
	params := map[string]interface{}{
		"model":       "text-davinci-003",
		"prompt":      "say hello to me 10 times",
		"max_tokens":  2048,
		"temperature": 0.9,
		"n":           1,
		"stream":      true,
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v", err)
	}

	msgs := []string{}

	for data := range res.Stream() {
		d := gjson.ParseBytes(data)
		s := d.Get("choices.0.text").String()
		msgs = append(msgs, s)

		log.Printf("%s\n", s)
	}

	fmt.Printf("message is: %s", strings.Join(msgs, ""))
	// Output: xxx
}
