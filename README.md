# Emoji Dict(ionary)
A constant for every emoji! Generated from the source of all emojis unicode.org!

## Usage
An emoji is a combination of one or more code points (runes). So the type is just a `[]rune` with a `String()` method.
```
fmt.Println(emojidict.RollingOnTheFloorLaughing)
// 🤣
fmt.Println(emojidict.RollingOnTheFloorLaughing.String())
// 🤣
```

### Generation
Make sure theres a line parser registered for whatever version of emoji you're pulling in `scraper/parser.go` and then do:
```
go generate
```
This pulls from the `unicode.org` ftp server, so dont use it a lot or they might shut that down, but I cache the request the first time so it's probably okay.
