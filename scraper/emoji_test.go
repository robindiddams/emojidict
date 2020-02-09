package scraper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLatestEmojiVersion(t *testing.T) {
	assert := assert.New(t)
	c := Client{}
	c.Connect()
	defer c.Close()
	val, err := c.GetLatestEmojiVersion()
	assert.NoError(err)
	fmt.Println("latest is", val)
	assert.NotEmpty(val)
}

// TODO: mock ftp and write real tests
