package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	assert := assert.New(t)
	trie := NewTrie()
	thermometer := []rune{0x1f321, 0xfe0f}
	violin := []rune{0x1f3bb}
	trie.Load(thermometer)
	trie.Load(violin)
	emojis := trie.Search("I love to play the "+string(thermometer)+"while i eat a "+string(violin), 2)
	assert.Len(emojis, 2)
}

func TestTrieLimit(t *testing.T) {
	assert := assert.New(t)
	trie := NewTrie()
	thermometer := []rune{0x1f321, 0xfe0f}
	violin := []rune{0x1f3bb}
	trie.Load(thermometer)
	trie.Load(violin)
	emojis := trie.Search("I love to play the "+string(thermometer)+"while i eat a "+string(violin), 1)
	assert.Len(emojis, 1)
	emojis = trie.Search("I love to play the "+string(thermometer)+"while i eat a "+string(violin), -1)
	assert.Len(emojis, 2)
}
