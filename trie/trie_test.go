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
	emojis := trie.Search("I love to play the "+string(thermometer)+" while i eat a "+string(violin), 2)
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

func TestZWJEmoji(t *testing.T) {
	assert := assert.New(t)
	trie := NewTrie()
	femaleSign := []rune{0x2640, 0xfe0f}
	// 129335, 8205, 9792, 65039
	womanShrugging := []rune{0x1f937, 0x200d, 0x2640, 0xfe0f}
	trie.Load(womanShrugging)
	trie.Load(femaleSign)
	var emojis [][]rune
	trie.SearchF("fix tests!! or did I "+string(womanShrugging), func(hit []rune, _ int) bool {
		emojis = append(emojis, hit)
		return false
	})
	assert.Len(emojis, 1)
	assert.Equal(womanShrugging, emojis[0])
}
