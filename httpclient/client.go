package httpclient

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type HttpOptions struct {
	SslVerify   bool
	ContentType string
	Timeout     int
}

var DefaultHttpOptions = &HttpOptions{
	SslVerify:   false,
	ContentType: "application/json",
	Timeout:     2000,
}

type HttpResponse struct {
	Data   []byte
	Status int
}

type _HttpClient struct {
	client *http.Client
}

func NewHttpClient(timeout int) HttpClient {
	return &_HttpClient{client: &http.Client{
		Timeout: time.Duration(time.Millisecond.Nanoseconds() * int64(timeout)),
	}}
}

//DoGet from url and return
func (hc *_HttpClient) DoGet(headers map[string]string, url string, options *HttpOptions) (*HttpResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("error in creating new request", err)
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if options.SslVerify == false {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	res, err := hc.client.Do(req)
	if err != nil {
		log.Println("Execute error ", err)
		return nil, err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Read error", err)
		return nil, err
	}

	return &HttpResponse{
		Data:   bytes,
		Status: res.StatusCode,
	}, nil
}
