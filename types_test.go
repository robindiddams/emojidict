package emojidict_test

import (
	"testing"

	"github.com/Robindiddams/emojidict"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(string(emojidict.RollingOnTheFloorLaughing), emojidict.RollingOnTheFloorLaughing.String())
}
