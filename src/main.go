package src

import (
	"time"
	"github.com/parnurzeal/gorequest"
	"fmt"
	"os"
)

const ApiVersion = 4

var debugEnabled = isDebugEnabled()

func isDebugEnabled() bool {
	return os.Getenv("GOCD_CLIENT_DEBUG") == "1"
}

// DefaultClient entrypoint for GoCD
type DefaultClient struct {
	Host    string `json:"host"`
	Request *gorequest.SuperAgent
}

// New GoCD Client
func New(host, username, password string) Client {
	client := DefaultClient{
		Host:    host,
		Request: gorequest.New().Timeout(60 * time.Second).SetBasicAuth(username, password),
	}
	return &client
}

func (c *DefaultClient) resolve(resource string) string {
	// TODO: Use a proper URL resolve to parse the string and append the resource

	if debugEnabled {
		fmt.Println("Api call to:", c.Host+resource)
	}

	return c.Host + resource
}
