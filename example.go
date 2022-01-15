package go_httpclient

import (
	"fmt"
	"io/ioutil"

	"github.com/CaptainStorm21/go-http-client/gohttp"
)

func exampleUsage() {

	client := gohttp.New()


	client.Get()

}
