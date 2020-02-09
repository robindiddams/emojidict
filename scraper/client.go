package scraper

import (
	"github.com/secsy/goftp"
)

// Client talks to unicode ftp server
type Client struct {
	ftpclient *goftp.Client
}

// Connect to the server
func (c *Client) Connect() (err error) {
	c.ftpclient, err = goftp.Dial("unicode.org")
	if err != nil {
		return
	}
	return
}

// Close the client
func (c *Client) Close() error {
	return c.ftpclient.Close()
}
