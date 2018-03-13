package gocd

import (
	"time"

	"github.com/parnurzeal/gorequest"
	"net/http"
	"crypto/tls"
)

// DefaultClient entrypoint for GoCD
type DefaultClient struct {
	Host    string `json:"host"`
	Request *gorequest.SuperAgent
}

// New GoCD Client
func New(host, username, password string) Client {

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout: 5 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}}

	superAgent := gorequest.New()
	superAgent.Client = httpClient

	client := DefaultClient{
		Host:    host,
		Request: gorequest.New().Timeout(60 * time.Second).SetBasicAuth(username, password),
	}
	return &client
}

func (c *DefaultClient) resolve(resource string) string {
	// TODO: Use a proper URL resolve to parse the string and append the resource
	return c.Host + resource
}

func (c *DefaultClient) SetHttpClientTransport(transport http.Transport) {
	c.Request.Client.Transport = &transport
}
