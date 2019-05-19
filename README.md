# http

[![Build Status](https://travis-ci.org/wujunze/http.svg)](https://travis-ci.org/wujunze/http)
[![codecov](https://codecov.io/gh/wujunze/http/branch/master/graph/badge.svg)](https://codecov.io/gh/wujunze/http)
[![Go Report Card](https://goreportcard.com/badge/github.com/wujunze/http)](https://goreportcard.com/report/github.com/wujunze/http)
[![GoDoc](https://godoc.org/github.com/wujunze/http?status.svg)](https://godoc.org/github.com/wujunze/http)

A simple http client for golang


## Usage

### Start using it

Download and install it:

```bash
$ go get github.com/wujunze/http
```

Import it in your code:

```go
import "github.com/wujunze/http"
```

### Example:

```go
package main

import (
	"github.com/wujunze/http"
)

func main() {
 
	params := make(map[string]string)
	params["name"] = "panda"
	curl := Curl{"https://github.com"}

	res, err := curl.Get("/", params)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(res))
}
```