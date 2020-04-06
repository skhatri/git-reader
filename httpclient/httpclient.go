package httpclient

type HttpClient interface {
	DoGet(headers map[string]string, url string, options *HttpOptions) (*HttpResponse, error)
}
