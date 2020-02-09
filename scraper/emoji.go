package scraper

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/blang/semver"
)

// EmojiFile is an expected filename in a given emoji's distribution
type EmojiFile string

const (
	// TestFileName is "emoji-test.txt", it contains every emoji
	TestFileName EmojiFile = "emoji-test.txt"
	// SequencesFileName is "emoji-sequences.txt", it contains ranges for all the emojis.
	// Useful for regex.
	SequencesFileName EmojiFile = "emoji-sequences.txt"
	// SequencesZWJFileName is "emoji-zwj-sequences.txt", it contains all the Zero Width Joiner emojis.
	// These are emojis composed of other emojis.
	SequencesZWJFileName EmojiFile = "emoji-zwj-sequences.txt"

	emojiPath = "Public/emoji"

	cacheDir = "dist"
)

// GetEmojiVersions will get every published version of emoji
func (c *Client) GetEmojiVersions() ([]string, error) {
	if c.ftpclient == nil {
		return nil, fmt.Errorf("client not connected")
	}
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

// GetLatestEmojiVersion gets the latest emoji version
func (c *Client) GetLatestEmojiVersion() (string, error) {
	versions, err := c.GetEmojiVersions()
	if err != nil {
		return "", err
	}
	return versions[len(versions)-1], nil
}

func makepath(version string, filename EmojiFile) string {
	return emojiPath + "/" + version + "/" + string(filename)
}

func fileExists(filename string) bool {
	f, err := filepath.Abs(filename)
	if err != nil {
		return false
	}
	_, err = os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func makeCachePath(path string) string {
	return filepath.Join(cacheDir, strings.ReplaceAll(path, "/", "_"))
}

func (c *Client) getFile(path string) ([]byte, error) {
	cacheFile := makeCachePath(path)
	// check cache
	if fileExists(cacheFile) {
		buf, err := ioutil.ReadFile(cacheFile)
		if err == nil {
			return buf, nil
		}
		fmt.Println("error reading cachefile:", err)
	}
	var buf bytes.Buffer
	if err := c.ftpclient.Retrieve(path, &buf); err != nil {
		return nil, err
	}
	data := buf.Bytes()
	// save to cache
	if !fileExists(cacheDir) {
		os.MkdirAll(cacheDir, 0777)
	}
	if err := ioutil.WriteFile(cacheFile, data, 0777); err != nil {
		fmt.Println("error writing cachefile:", err)
	}
	return data, nil
}

// GetEmojiFile returns file for version
func (c *Client) GetEmojiFile(version string, file EmojiFile) ([]byte, error) {
	return c.getFile(makepath(version, TestFileName))
}

// GetLatestEmojiFile is a convenience for creating a client, connectecting, finding the latest emoji version,
// and getting a file from it.
func GetLatestEmojiFile(file EmojiFile) ([]byte, string, error) {
	c := Client{}
	if err := c.Connect(); err != nil {
		return nil, "", fmt.Errorf("error connecting: %w", err)
	}
	defer c.Close()
	v, err := c.GetLatestEmojiVersion()
	if err != nil {
		return nil, "", fmt.Errorf("error getting latest emoji version: %w", err)
	}
	buf, err := c.getFile(makepath(v, file))
	return buf, v, err
}
