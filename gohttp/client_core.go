package gohttp

import (
	"io/ioutil"
	"net/http"
)

func (c *httpClient) do(
	method string,
	url string,
	headers http.Header, 
	body interface{},
) (*http.Response, error) {

	client := http.Client{}

	request, err := http.NewRequest(httpMethod, url, body:nil)

	if err != nil{
		return nil, errors.New(text: "unable to render")
	}
	return client.Do(request)
	
}