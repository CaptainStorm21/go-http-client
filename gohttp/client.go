package gohttp

type httpClient struct {
}

func New() HttpClient {
	client := &httpClient{}
	return client
}

type HttpClient interface {
	Get()
	Post()
	Put()
	Patch()
	Delete()
}

// GET
func (c *httpClient) Get() {

}

// POST
func (c *httpClient) Post() {

}

// PUT
func (c *httpClient) Put() {

}

// PATCH
func (c *httpClient) Patch() {

}

// DELETE
func (c *httpClient) Delete() {

}
