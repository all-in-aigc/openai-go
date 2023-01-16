package gpt3

import (
	"fmt"
	"time"

	"github.com/idoubi/goutils/request"
)

// Client: GPT-3 client
type Client struct {
	opts          *Options // custom options
	requestClient *request.Client
}

// NewClient: new GPT-3 Client
func NewClient(opts *Options) (*Client, error) {
	// set default options
	if opts.Timeout <= 0 {
		opts.Timeout = 30 * time.Second
	}

	cli := &Client{opts: opts}

	// set request client
	cli.requestClient = request.NewClient(&request.Options{
		BaseUri: "https://api.openai.com",
		Debug:   opts.Debug,
		Timeout: opts.Timeout,
	})

	return cli, nil
}

// Get: request api with get method
func (cli *Client) Get(uri string, args ...map[string]interface{}) (*request.Result, error) {
	params, headers := cli.parseArgs(args...)

	return cli.getRequestClient().Get(uri, params, headers)
}

// Post: request api with post method
func (cli *Client) Post(uri string, args ...map[string]interface{}) (*request.Result, error) {
	data, headers := cli.parseArgs(args...)

	return cli.getRequestClient().Post(uri, data, headers)
}

// parseArgs: parse request args and append api_key
func (cli *Client) parseArgs(args ...map[string]interface{}) (map[string]interface{}, map[string]interface{}) {
	params := map[string]interface{}{}
	headers := map[string]interface{}{}

	if len(args) > 0 {
		params = args[0]
	}
	if len(args) > 1 {
		headers = args[1]
	}
	headers["Authorization"] = fmt.Sprintf("Bearer %s", cli.opts.ApiKey)

	return params, headers
}

// getRequestClient: get request handler
func (cli *Client) getRequestClient() *request.Client {
	return cli.requestClient
}
