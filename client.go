package main

import (
	"github.com/modouwifi/md/httpkit"
)

type Client struct {
	ApiURL    string
	UserAgent string
	Request   *httpkit.Request
}

func (c *Client) NewRequest(method, uri string) (*httpkit.Request, error) {
	c.Request = httpkit.New(method, c.ApiURL+uri, c.UserAgent)
	c.Request.SetProtocolVersion("HTTP/1.1")
	return c.Request, nil
}

func NewClient(url, userAgent string) *Client {
	return &Client{
		ApiURL:    url,
		UserAgent: userAgent,
	}
}
