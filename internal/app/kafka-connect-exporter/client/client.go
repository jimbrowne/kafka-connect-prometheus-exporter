package client

import (
	"net/http"
	"time"
)

type AuthCredentials struct {
	User     string
	Password string
}

type client struct {
	client          http.Client
	baseUrl         string
	authCredentials *AuthCredentials
}

func NewClient(baseUrl string, authCredentials *AuthCredentials) Client {
	return &client{
		client: http.Client{
			Timeout: 3 * time.Second,
		},
		baseUrl:         baseUrl,
		authCredentials: authCredentials,
	}
}
