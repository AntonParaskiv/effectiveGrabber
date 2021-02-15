package HttpClientInterface

type Client interface {
	Get(url string) (responseBody []byte, statusCode int, err error)
}
