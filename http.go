package http

import (
	"net/url"
)

type Client interface {
	Get(uri string, params map[string]string) ([]byte, error)
	Post(uri string, params url.Values) (string, error)
}
