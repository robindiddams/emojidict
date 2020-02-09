package emojidict

//go:generate go run gen.go
//go:generate go fmt emojis.go
//go:generate go fmt categories.go

// Emoji is a set of code points
type Emoji []rune

func (e Emoji) String() string {
	return string(e)
}
