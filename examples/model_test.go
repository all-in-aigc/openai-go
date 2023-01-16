package examples

import (
	"fmt"
	"log"
)

func ExampleListModels() {
	cli := getClient()

	uri := "/v1/models"

	res, err := cli.Get(uri)

	if err != nil {
		log.Fatalf("request api failed: %v", err)
	}

	for _, v := range res.Get("data").Array() {
		fmt.Printf("model id: %s\n", v.Get("id").String())
	}

	// Output: xxx
}

func ExampleRetrieveModel() {
	cli := getClient()

	model := "text-davinci-003"

	uri := fmt.Sprintf("/v1/models/%s", model)

	res, err := cli.Get(uri)

	if err != nil {
		log.Fatalf("request api failed: %v", err)
	}

	fmt.Println(res.GetString("id"))

	// Output: xxx
}
