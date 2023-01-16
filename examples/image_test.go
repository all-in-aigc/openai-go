package examples

import (
	"fmt"
	"log"
)

func ExampleCreateImage() {
	cli := getClient()

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

	// Output: xxx
}
