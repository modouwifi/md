package mdclient

import (
	"github.com/fundon/md/httprequest"
)

type Client struct {
	ApiURL    string
	UserAgent string
	Request   *httprequest.Request
}

func (c *Client) NewRequest(method, uri string) (*httprequest.Request, error) {
	c.Request = httprequest.New(method, c.ApiURL+uri, c.UserAgent)
	c.Request.SetProtocolVersion("HTTP/1.1")
	return c.Request, nil
}

func New(url, userAgent string) *Client {
	return &Client{
		ApiURL:    url,
		UserAgent: userAgent,
	}
}
