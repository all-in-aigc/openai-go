package examples

import (
	"fmt"
	"log"

	"github.com/tidwall/gjson"
)

func ExampleFunctionCallArgs() {
	cli := getClient()

	userQuestion := "What is the weather like in Boston?"

	uri := "/v1/chat/completions"
	params := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": userQuestion,
			},
		},
		"functions": getFuncs(),
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v", err)
	}

	funcName := res.Get("choices.0.message.function_call.name").String()
	funcArgs := res.Get("choices.0.message.function_call.arguments").String()

	if funcName == "" || funcArgs == "" {
		log.Fatalf("function call get args failed: %s", res)
	}

	fmt.Printf("function call name: %s, args: %v", funcName, funcArgs)
	// Output: xxx
}

func ExampleFunctionCallArgsWithStream() {
	cli := getClient()

	userQuestion := "What is the weather like in Boston?"

	uri := "/v1/chat/completions"
	params := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": userQuestion,
			},
		},
		"functions": getFuncs(),
		"stream":    true,
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v", err)
	}

	var funcName string
	var funcArgs string

	for data := range res.Stream() {
		d := gjson.ParseBytes(data)

		fc := d.Get("choices.0.delta.function_call").String()
		if fc != "" {
			if name := d.Get("choices.0.delta.function_call.name").String(); name != "" {
				funcName = name
			}
			funcArgs += d.Get("choices.0.delta.function_call.arguments").String()

			continue
		}
		log.Printf("%s\n", d)
	}

	fmt.Printf("function call name: %s, args: %v", funcName, funcArgs)
	// Output: xxx
}

func ExampleFunctionCallSummarize() {
	cli := getClient()

	userQuestion := "What is the weather like in Boston?"

	// value get from function call args
	funcName := "get_current_weather"
	funcArgs := "{\n  \"location\": \"Boston, MA\"\n}"

	// value get from api call
	apiRes := ""

	uri := "/v1/chat/completions"
	params := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": userQuestion,
			},
			{
				"role":    "assistant",
				"content": nil,
				"function_call": map[string]interface{}{
					"name":      funcName,
					"arguments": funcArgs,
				},
			},
			{
				"role":    "function",
				"name":    funcName,
				"content": apiRes,
			},
		},
		"functions": getFuncs(),
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v", err)
	}

	message := res.Get("choices.0.message.content").String()
	fmt.Printf("message is: %s", message)
	// Output: xxx
}

func getFuncs() []map[string]interface{} {
	funcs := []map[string]interface{}{}

	getWeatherFunc := map[string]interface{}{
		"name":        "get_current_weather",
		"description": "Get the current weather in a given location",
		"parameters": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"location": map[string]interface{}{
					"type":        "string",
					"description": "The city and state, e.g. San Francisco, CA",
				},
				"unit": map[string]interface{}{
					"type": "string",
					"enum": []string{"celsius", "fahrenheit"},
				},
			},
			"required": []string{"location"},
		},
	}

	funcs = append(funcs, getWeatherFunc)

	return funcs
}
