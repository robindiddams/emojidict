package emojidict

import (
	"github.com/robindiddams/go-trie"
)

// Parser is a wrapper on a trie
type Parser struct {
	emojtree trie.Trie
}

// NewEmojiParser returns an emoji parser
func NewEmojiParser() Parser {
	t := trie.NewTrie()
	for _, emoji := range All {
		t.Load(emoji)
	}
	return Parser{t}
}

// FindEmojis returns all emojis in s
func (p *Parser) FindEmojis(s string) []Emoji {
	var emojis []Emoji
	p.emojtree.SearchF(s, func(hit []rune, _ int) bool {
		emojis = append(emojis, Emoji(hit))
		return false
	})
	return emojis
}

// RemoveEmojis finds all emojis in a string and removes them,
// this uses the trie which is less efficient than if the emoji
// ranges were loaded into a regex.
func (p *Parser) RemoveEmojis(s string) string {
	str := []rune(s)
	var place int
	p.emojtree.SearchF(s, func(hit []rune, at int) bool {
		at = at - place
		str = append(str[:at], str[at+len(hit):]...)
		place += len(hit)
		return false
	})
	return string(str)
}
