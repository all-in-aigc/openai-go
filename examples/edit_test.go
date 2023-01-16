package examples

import (
	"fmt"
	"log"
)

func ExampleCreateEdit() {
	cli := getClient()

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

	// Output: xxx
}
