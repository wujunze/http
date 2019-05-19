package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Curl http请求类
type Curl struct {
	Domain string
}

// Get http-get请求
func (c Curl) Get(uri string, params map[string]string) ([]byte, error) {
	var requestURL string
	if uri == "" {
		requestURL = c.Domain
	} else {
		requestURL = c.Domain + "/" + uri
	}
	if len(params) > 0 {
		parseURL, parseErr := url.Parse(requestURL)
		if parseErr != nil {
			return []byte{}, parseErr
		}
		q := parseURL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		parseURL.RawQuery = q.Encode()
		requestURL = parseURL.String()
	}
	result, err := http.Get(requestURL)
	if err != nil {
		fmt.Print( "curl_get", "get:"+requestURL+" error,errMessage:"+err.Error(), "")
		return []byte{}, err
	}
	resultByte, err1 := ioutil.ReadAll(result.Body)
	defer result.Body.Close()
	if err1 != nil {
		fmt.Print("curl_ioutil_read", "get:"+requestURL+" error,errMessage:"+err1.Error(), "")
		return []byte{}, err1
	}
	return resultByte, nil
}

// Post http的post请求
func (c Curl) Post(uri string, params url.Values) (string,error) {
	requestURL := c.Domain + "/" + uri
	result, err := http.PostForm(requestURL, params)
	if err != nil {
		fmt.Print("curl_post", "post:"+requestURL+" error,errMessage:"+err.Error(), "")
		return "",err
	}
	defer result.Body.Close()
	resultByte, err1 := ioutil.ReadAll(result.Body)
	if err1 != nil {
		return "",err1
	}
	return string(resultByte),nil
}