package gohttp

import "net/http"

type httpClient struct {
}

func New() HttpClient {
	client := &httpClient{}
	return client
}

type HttpClient interface {
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header)(*http.Response, error)
}

// GET
func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, body:nil)
}

// POST
func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error){
		return c.do(http.MethodPost, url, headers, body)

}

// PUT
func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error){
	return c.do(http.MethodPut, url, headers, body: nil)
}

// PATCH
func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error){
	return c.do(http.MethodPatch, url, headers, body: nil)
}

// DELETE
func (c *httpClient) Delete(url string, headers http.Header)(*http.Response, error) {
		return c.do(http.MethodDelete, url, headers, body: nil)

}
