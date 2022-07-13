package client

import (
	"github.com/prometheus/common/log"
	"net/http"
)

func (c *client) Get(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.baseUrl+path, nil)
	if err != nil {
		log.Errorf("failed creating request: %v", err)
		return nil, err
	}
	if c.authCredentials != nil {
		req.SetBasicAuth(c.authCredentials.User, c.authCredentials.Password)
	}
	return c.client.Do(req)
}
