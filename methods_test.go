package emojidict_test

import (
	"testing"

	"github.com/Robindiddams/emojidict"

	"github.com/stretchr/testify/assert"
)

func TestEmojiParser(t *testing.T) {
	assert := assert.New(t)
	parser := emojidict.NewEmojiParser()
	emojis := parser.FindEmojis("fix tests ✅")
	assert.Len(emojis, 1)
}

func TestParserWithHumanEmoji(t *testing.T) {
	assert := assert.New(t)
	parser := emojidict.NewEmojiParser()
	emojis := parser.FindEmojis(emojidict.WomanShrugging.String() + "  ")
	assert.Len(emojis, 1)
	assert.EqualValues(emojidict.WomanShrugging, emojis[0])
}

func TestRemoveEmojis(t *testing.T) {
	assert := assert.New(t)
	parser := emojidict.NewEmojiParser()
	cleanStr := parser.RemoveEmojis("fix tests ✅!!")
	assert.EqualValues("fix tests !!", cleanStr)
}

func TestRemoveMultipleEmojis(t *testing.T) {
	assert := assert.New(t)
	parser := emojidict.NewEmojiParser()
	cleanStr := parser.RemoveEmojis("fix tests ✅!!🤔 or did I " + emojidict.WomanShrugging.String())
	assert.EqualValues("fix tests !! or did I ", cleanStr)
}
