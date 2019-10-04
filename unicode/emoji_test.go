package unicode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineToCharacter(t *testing.T) {
	assert := assert.New(t)
	char := lineToCharacter("1F600                                      ; fully-qualified     # 😀 grinning face")
	assert.NotNil(char)
	assert.EqualValues("grinning face", char.Description)
	assert.EqualValues("😀", char.Value)
}
