package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	host string
}

func NewClient(port uint32) *Client {
	return &Client{
		host: fmt.Sprintf("localhost:%d",
			port),
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
	return result, httpGet(url.String(), &result)
}

// httpGet sends a http GET request and decodes the JSON response in
// provided output.
func httpGet(url string, output interface{}) error {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	request.Header.Add("Accept", "application/json")
	client := http.Client{}
	response, err := client.Do(request)
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
