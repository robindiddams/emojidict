package unicode

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
		return err
	}
	return nil
}

// Close the client
func (c *Client) Close() (err error) {
	return c.ftpclient.Close()
}

// Character is a single unicode character
type Character struct {
	Value       string
	Description string
	Hex         string
}
