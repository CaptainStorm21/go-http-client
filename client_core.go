package gohttp

import (
	// "io/ioutil"
	"net/http"
	"errors"
)

func (c *httpClient) do (method string, url string, headers http.Header, body interface{}) (*http.Response, error) {

	client := http.Client{}

	request, err := http.NewRequest(method, url, body:nil)

	if err != nil {
		return nil, errors.New(text "unable to process a request")
	}

	return  client.Do(request)
	// if err !=nil {
	// 	panic (err)
	// }

	// return response, err

	// defer response.Body.Close()
	// bytes, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }


}