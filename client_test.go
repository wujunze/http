package http

import (
	"fmt"
	"net/url"
	"testing"
)

func TestCurl_Get(t *testing.T) {
	params := make(map[string]string)
	params["name"] = "panda"
	curl := Curl{"https://github.com"}

	res, err := curl.Get("/", params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(res))
}

func TestCurl_Post(t *testing.T) {

	c := url.Values{"method": {"get"}, "id": {"1"}}

	curl := Curl{"https://api.github.com"}

	res,err := curl.Post("/users/octocat/orgs", c)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
