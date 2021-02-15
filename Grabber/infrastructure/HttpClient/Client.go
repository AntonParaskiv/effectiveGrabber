package HttpClient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	baseUrl url.URL
}

func New() (s *Client) {
	s = new(Client)
	return
}

func (c *Client) SetBaseUrl(baseUrl url.URL) *Client {
	c.baseUrl = baseUrl
	return c
}

func (c *Client) Get(additionalUrl string) (responseBody []byte, statusCode int, err error) {
	// make full url
	fullUrl := c.baseUrl
	fullUrl.Path = path.Join(c.baseUrl.Path, additionalUrl)

	// request
	response, err := http.Get(fullUrl.String())
	if err != nil {
		err = fmt.Errorf("http get failed: %w", err)
		return
	}
	defer response.Body.Close()

	// get response
	statusCode = response.StatusCode
	responseBody, err = ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("read response body failed: %w", err)
		return
	}

	return
}
