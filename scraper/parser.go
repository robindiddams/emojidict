package scraper

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/serenize/snaker"
)

// EmojiLine is the name part and hex part
type EmojiLine struct {
	Name       string
	ProperName string
	CodePoints []string
}

// Parser parses one line of an emoji file
type Parser func(string) *EmojiLine

var parsers map[string]Parser

// ParseEmojis will parse a line
func ParseEmojis(version string, file EmojiFile, line string) (*EmojiLine, error) {
	fn := parsers[makeParserKey(version, file)]
	if fn == nil {
		return nil, fmt.Errorf("no parser registered for version %s, file %s", version, file)
	}
	l := fn(line)
	if l != nil {
		l.clean()
	}
	return l, nil
}

func makeParserKey(version string, file EmojiFile) string {
	return version + string(file)
}

// RegisterParserFunc lets one register a parser for a file/version
func RegisterParserFunc(version string, file EmojiFile, fn Parser) {
	if parsers == nil {
		parsers = make(map[string]Parser)
	}
	parsers[makeParserKey(version, file)] = fn
}

// Clean makes all the codepoints lowercase, removes spaces from names
func (l *EmojiLine) clean() {
	// make sure they're all lowercass
	for i, codePoint := range l.CodePoints {
		l.CodePoints[i] = strings.ToLower(codePoint)
	}
	propername := strings.TrimSpace(l.Name)
	propername = strings.ToLower(propername)
	// remove bad characters
	propername = strings.ReplaceAll(propername, ":", "")
	propername = strings.ReplaceAll(propername, ";", "")
	propername = strings.ReplaceAll(propername, "&", "")
	propername = strings.ReplaceAll(propername, "!", "")
	propername = strings.ReplaceAll(propername, ".", "")
	propername = strings.ReplaceAll(propername, ",", "")
	propername = strings.ReplaceAll(propername, "`", "")
	propername = strings.ReplaceAll(propername, "'", "")
	propername = strings.ReplaceAll(propername, "“", "")
	propername = strings.ReplaceAll(propername, "”", "")
	propername = strings.ReplaceAll(propername, "’", "")
	propername = strings.ReplaceAll(propername, "\"", "")
	// fix bad prefixes
	replacePrefixWith := func(old, new string) {
		if strings.HasPrefix(propername, old) {
			propername = strings.Replace(propername, old, new, 1)
		}
	}
	replacePrefixWith("1st", "first")
	replacePrefixWith("2nd", "second")
	replacePrefixWith("3rd", "third")

	// snake-case-ify
	propername = strings.ReplaceAll(propername, " ", "_")
	propername = strings.ReplaceAll(propername, "-", "_")
	propername = strings.ReplaceAll(propername, "(", "_")
	propername = strings.ReplaceAll(propername, ")", "_")
	// convert!
	l.ProperName = snaker.SnakeToCamel(propername)
}

// testFileV13 parses "emoji-test.txt" for emoji version 13
func testFileV13(line string) *EmojiLine {
	rejectex := regexp.MustCompile("^[#0]")
	acceptex := regexp.MustCompile("(.*)\\s+;\\sfully-qualified\\s+#.*E\\d+\\.\\d\\s(.*)")
	if !rejectex.MatchString(line) && acceptex.MatchString(line) {
		matches := acceptex.FindStringSubmatch(line)
		return &EmojiLine{
			CodePoints: strings.Split(strings.TrimSpace(matches[1]), " "),
			Name:       strings.TrimSpace(matches[2]),
		}
	}
	return nil
}

func init() {
	RegisterParserFunc("13.0", TestFileName, testFileV13)
}
