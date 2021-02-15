package ApiRepository

import (
	"encoding/json"
	"fmt"
)

func (r *Repository) requestJsonGet(url string, response interface{}, wantStatusCode int) (err error) {
	// send request
	responseBytes, statusCode, err := r.httpClient.Get(url)
	if err != nil {
		err = fmt.Errorf("http client get failed: %w", err)
		return
	}

	// check status code
	if statusCode != wantStatusCode {
		err = fmt.Errorf("http client get bad status: %d", statusCode)
		return
	}

	// unmarshal response
	err = json.Unmarshal(responseBytes, response)
	if err != nil {
		err = fmt.Errorf("unmarshal response failed: %w", err)
		return
	}

	return
}
