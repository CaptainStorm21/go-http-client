package main

import (
	"fmt"
	"io/ioutil"

	"github.com/CaptainStorm21/go-http-client/gohttp"
)

func main() {

	client := gohttp.New()

	response, err := client.Get(url: "https://api.github.com", headers: nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode) 
}
