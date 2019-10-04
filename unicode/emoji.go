package unicode

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/blang/semver"
)

const emojiPath = "Public/emoji"

// GetEmojiVersions will get every published version of emoji
func (c *Client) GetEmojiVersions() ([]string, error) {
	files, err := c.ftpclient.ReadDir(emojiPath)
	if err != nil {
		return nil, err
	}
	var versions semver.Versions
	for _, file := range files {
		if file.IsDir() {
			ver, err := semver.Parse(file.Name() + ".0")
			if err != nil {
				return nil, err
			}
			versions = append(versions, ver)
		}
	}
	sort.Sort(versions)
	var out []string
	for _, ver := range versions {
		out = append(out, fmt.Sprintf("%d.%d", ver.Major, ver.Minor))
	}
	return out, nil
}

func makeEmojiTestPath(version string) string {
	return emojiPath + "/" + version + "/emoji-test.txt"
}

func (c *Client) getEmojiTest(version string) (string, error) {
	var buf bytes.Buffer
	if err := c.ftpclient.Retrieve(makeEmojiTestPath(version), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func lineToCharacter(line string) *Character {
	rejectex := regexp.MustCompile("^[#0]")
	acceptex := regexp.MustCompile("(.*)\\s+;\\sfully-qualified\\s+#\\s([^\\s]*)\\s(.*)")
	if !rejectex.MatchString(line) && acceptex.MatchString(line) {
		matches := acceptex.FindStringSubmatch(line)
		return &Character{
			Value:       matches[2],
			Description: strings.TrimSpace(matches[3]),
			Hex:         strings.TrimSpace(matches[1]),
		}
	}
	return nil
}

// GetEmojis gets all the emojis for a given character
func (c *Client) GetEmojis(version string) ([]Character, error) {
	test, err := c.getEmojiTest(version)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(test, "\n")
	var chars []Character
	for _, line := range lines {
		if char := lineToCharacter(line); char != nil {
			chars = append(chars, *char)
		}
	}

	return chars, nil
}
