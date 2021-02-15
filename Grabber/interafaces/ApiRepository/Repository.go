package ApiRepository

import "github.com/AntonParaskiv/effectiveGrabber/Grabber/interafaces/HttpClientInterface"

const (
	endpointOne = "/api/one"
	endpointTwo = "/api/two"
)

type Repository struct {
	httpClient HttpClientInterface.Client
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetHttpClient(httpClient HttpClientInterface.Client) *Repository {
	r.httpClient = httpClient
	return r
}
