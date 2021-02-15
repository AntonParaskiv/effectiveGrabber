package ApiRepository

import (
	"fmt"
	"net/http"
)

type Response struct {
	Result int
}

func (r *Repository) GetOne() (err error) {
	err = r.getAny(endpointOne)
	return
}

func (r *Repository) GetTwo() (err error) {
	err = r.getAny(endpointTwo)
	return
}

func (r *Repository) getAny(endpoint string) (err error) {
	response := new(Response)

	// request
	err = r.requestJsonGet(endpoint, response, http.StatusOK)
	if err != nil {
		err = fmt.Errorf("request failed: %w", err)
		return
	}

	// check result
	if response.Result != 1 {
		err = fmt.Errorf("invalid response result: %d", response.Result)
		return
	}

	return
}
