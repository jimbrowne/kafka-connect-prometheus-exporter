package client

import "net/http"

type Client interface {
	Get(path string) (*http.Response, error)
}
