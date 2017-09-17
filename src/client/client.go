package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	host   string
	client *http.Client
}

func NewClient(port uint32, maxIdleConnsPerHost int) *Client {
	transport := &http.Transport{
		MaxIdleConnsPerHost: maxIdleConnsPerHost,
	}
	return &Client{
		host: fmt.Sprintf("localhost:%d",
			port),
		client: &http.Client{Transport: transport},
	}
}

func (c *Client) FizzBuzz(int1, int2, limit int, string1, string2 string) (
	[]string, error) {

	result := []string{}
	url, err := url.Parse(fmt.Sprintf("http://%s/fizzbuzz", c.host))
	if err != nil {
		return result, err
	}
	values := url.Query()
	values.Add("int1", fmt.Sprintf("%v", int1))
	values.Add("int2", fmt.Sprintf("%v", int2))
	values.Add("limit", fmt.Sprintf("%v", limit))
	values.Add("string1", string1)
	values.Add("string2", string2)
	url.RawQuery = values.Encode()
	return result, c.httpGet(url.String(), &result)
}

// httpGet sends a http GET request and decodes the JSON response in
// provided output.
func (c *Client) httpGet(url string, output interface{}) error {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	request.Header.Add("Accept", "application/json")
	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("server request failed, error code: %d",
			response.StatusCode)
	}
	return json.NewDecoder(response.Body).Decode(output)
}
